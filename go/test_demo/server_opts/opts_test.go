package main

import (
  "testing"
  "reflect"
)

func TestWithTls(t *testing.T) {
  if serv := newServer(WithTls); !serv.Opts.tls {
    t.Error("WithTls expected be true, but false got")
  }
}

func TestWithMaxConn(t *testing.T) {
  if serv := newServer(WithMaxConn(100)); serv.Opts.maxConn != 100 {
    t.Errorf("WithMaxConn expected be 100, but %d got", serv.Opts.maxConn)
  }
}

func TestWithId(t *testing.T) {
  if serv := newServer(WithId("testing")); serv.Opts.id != "testing" {
    t.Errorf("WithId expected be 'testing', but %s got", serv.Opts.id)
  }
}

func TestMul(t *testing.T) {
  cases := []struct {
    Name string
    WithVal, Expected interface{}
  }{
    {"maxConn", 120, 120},
    {"id", "testing2", "testing2"},
    {"tls", true, true},
  }

  serv := newServer(
    WithMaxConn(cases[0].WithVal.(int)),
    WithId(cases[1].WithVal.(string)), 
    WithTls)

  optsType := reflect.TypeOf(serv.Opts)
  optsValue := reflect.ValueOf(serv.Opts)

  for index, c := range cases {
    t.Run(c.Name, func(t *testing.T) {
      fieldType := optsType.Field(index) 
      field := optsValue.Field(index)

      var value interface{}

      switch field.Kind() {
      case reflect.String:
        value = field.String()
      case reflect.Int:
        value = field.Int()
      case reflect.Bool:
        value = field.Bool()
      }

      if fieldType.Name != cases[index].Name {
        t.Fatalf(
          "field name: %+v, but %+v got", 
          cases[index].Name, fieldType.Name)
      } 

      if cases[index].Expected != value {
        t.Fatalf(
          "field name: %s, val: %+v, but %+v got", 
          fieldType.Name, cases[index].Expected, value)
      } 
    })
  }
}
