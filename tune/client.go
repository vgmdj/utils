package tune

import (
	"context"
	"errors"
	"os"
	"time"

	"go.uber.org/zap/zapcore"

	"go.uber.org/zap"
)

// Logger logger interface
type Logger interface {
	Info(args ...interface{})
}

// DefaultLogger use default logger ,if not set
var DefaultLogger *zap.SugaredLogger

func init() {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(os.Stderr),
		zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= zapcore.InfoLevel
		}),
	)

	DefaultLogger = zap.New(core,
		zap.WithCaller(true),
	).Sugar()
}

// Resize the handler of resize
type Resize func(ctx context.Context, observation, toSize int64) error

// Client tune's client
type Client struct {
	collector Collector // collector client
	resize    Resize    // resize handler
	ac        Tune      // Arithmetic Unit
	logger    Logger    // logger
}

// NewClient return the object of client
func NewClient(options ...Option) *Client {
	cli := &Client{
		logger: DefaultLogger,
	}

	for _, option := range options {
		option(cli)
	}

	err := cli.CheckConfig()
	if err != nil {
		panic(err.Error())
	}

	return cli
}

// CheckConfig check the config is valid or not
func (cli *Client) CheckConfig() error {
	if cli == nil {
		return errors.New("client can not be nil")
	}

	if cli.ac == nil {
		return errors.New("ac can not be nil")
	}

	if cli.resize == nil {
		return errors.New("resize func can not be nil")
	}

	return nil

}

// Option option func
type Option func(cli *Client)

// WithResize set the resize config
func WithResize(resize Resize) Option {
	return func(cli *Client) {
		cli.resize = resize
	}
}

func WithCollector(c Collector) Option {
	return func(cli *Client) {
		cli.collector = c
	}
}

// WithArithmeticUnit set the ac config
func WithArithmeticUnit(t Tune) Option {
	return func(cli *Client) {
		cli.ac = t
	}
}

// WithLogger set the logger config
func WithLogger(logger Logger) Option {
	return func(cli *Client) {
		cli.logger = logger
	}
}

// ListenAndResize keep resizing
func (cli *Client) ListenAndResize(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			cli.logger.Info("Tune Exit by Context Done")
			return

		default:
			err := cli.exec(ctx)
			if err != nil {
				cli.logger.Info("Tune Exec Failed", "error", err.Error())
			}
			time.Sleep(time.Second)

		}

	}

}

// exec get collector's avg and resize to suitable size
func (cli *Client) exec(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, time.Minute*5)
	defer cancel()
	now := int64(cli.collector.Avg())
	to := cli.ac.SuitableSize(now)
	err := cli.resize(ctx, now, to)
	if err != nil {
		return err
	}

	return nil
}
