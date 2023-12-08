package main

type Server struct {
  Opts
}

type Opts struct {
  maxConn int
  id string
  tls bool
}

type OptsFunc func(*Opts)

func defaultOpt() Opts {
  return Opts{
    maxConn: 10,
    id: "default",
    tls: false,
  }

}

func WithMaxConn(maxConn int) OptsFunc {
  return func(o *Opts) {
    o.maxConn = maxConn
  }
}

func WithId(id string) OptsFunc {
  return func(o *Opts) {
    o.id = id
  }
}

func WithTls(o *Opts)  {
  o.tls = true 
}

func newServer(optsFuncs ...OptsFunc) *Server {
  opt := defaultOpt()
  for _, f := range optsFuncs {
    f(&opt)
  }

  return &Server {
    Opts: opt,
  }
}

