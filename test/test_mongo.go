package test

import (
	"github.com/zeromicro/go-zero/core/stores/mon"
	"testing"
)

func MongoTest(t *testing.T) {
	conn := mon.MustNewModel("mongodb://localhost:27017", "test", "users")
	defer conn.Clone()

}
