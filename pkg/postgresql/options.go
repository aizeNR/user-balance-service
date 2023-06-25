package postgresql

// Option -.
type Option func(*Postgres)

// WithMaxPoolSize -.
func WithMaxPoolSize(size int) Option {
	return func(c *Postgres) {
		c.maxPoolSize = size
	}
}

func WithTracer(tracer tracer) Option {
	return func(c *Postgres) {
		c.tracer = tracer
	}
}
