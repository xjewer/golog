package golog_test

import (
	"bytes"
	"testing"

	"github.com/chapsuk/golog"
	test "github.com/smartystreets/goconvey/convey"
)

// FakeFormatter for test
type FakeFormatter struct{}

// Format wirte to buffer only log message
func (f *FakeFormatter) Format(b *bytes.Buffer, lvl golog.Level, ctx golog.Context, msg string, trace []byte) *bytes.Buffer {
	b.WriteString(msg)
	return b
}

func TestExportGologPackage(t *testing.T) {
	test.Convey("Check package methods exported and work fine", t, func() {
		output := &FakeWriter{}
		golog.SetOutput(output)
		golog.SetLevel(golog.DebugLevel)

		test.Convey("check setFormatter, setOutput and SetLevel methods", func() {
			golog.SetFormatter(&FakeFormatter{})
			golog.SetLevel(golog.ErrorLevel)

			golog.Error(foo)
			golog.Info(bar)
			test.So(output.log, test.ShouldHaveLength, 1)
			test.So(output.log, test.ShouldContain, foo)
		})

		test.Convey("check with context", func() {
			ctx := golog.Context{
				"foo": foo,
				"bar": bar,
			}
			log := golog.WithContext(ctx)

			test.So(log.GetContext(), test.ShouldResemble, ctx)
			test.So(golog.GetContext(), test.ShouldResemble, golog.Context{})
		})

		test.Convey("check Print* methods", func() {
			golog.Print(foo)
			golog.Printf("%s", bar)
			golog.Println(foo, bar)

			test.So(output.log, test.ShouldHaveLength, 3)
		})

		test.Convey("check Debug* methods", func() {
			golog.Debug(foo)
			golog.Debugf("%s", bar)
			golog.DebugCtx(golog.Context{}, baz)
			golog.DebugfCtx(golog.Context{}, "%s", bee)
			golog.Debugln(foo, bar)

			test.So(output.log, test.ShouldHaveLength, 5)
		})

		test.Convey("check Info* methods", func() {
			golog.Info(foo)
			golog.Infof("%s", bar)
			golog.InfoCtx(golog.Context{}, baz)
			golog.InfofCtx(golog.Context{}, "%s", bee)
			golog.Infoln(foo, bar)

			test.So(output.log, test.ShouldHaveLength, 5)
		})

		test.Convey("check Warn* methods", func() {
			golog.Warn(foo)
			golog.Warnf("%s", bar)
			golog.WarnCtx(golog.Context{}, baz)
			golog.WarnfCtx(golog.Context{}, "%s", bee)
			golog.Warnln(foo, bar)

			test.So(output.log, test.ShouldHaveLength, 5)
		})

		test.Convey("check Error* methods", func() {
			golog.Error(foo)
			golog.Errorf("%s", bar)
			golog.ErrorCtx(golog.Context{}, baz)
			golog.ErrorfCtx(golog.Context{}, "%s", bee)
			golog.Errorln(foo, bar)

			test.So(output.log, test.ShouldHaveLength, 5)
		})

		test.Convey("check Fatal* methods", func() {
			golog.Fatal(foo)
			golog.Fatalf("%s", bar)
			golog.FatalCtx(golog.Context{}, baz)
			golog.FatalfCtx(golog.Context{}, "%s", bee)
			golog.Fatalln(foo, bar)

			test.So(output.log, test.ShouldHaveLength, 5)
		})

		test.Convey("check Panic* methods", func() {
			test.So(func() {
				golog.Panic(foo)
			}, test.ShouldPanic)
			test.So(func() {
				golog.Panicf("%s", bar)
			}, test.ShouldPanic)
			test.So(func() {
				golog.PanicCtx(golog.Context{}, baz)
			}, test.ShouldPanic)
			test.So(func() {
				golog.PanicfCtx(golog.Context{}, "%s", bee)
			}, test.ShouldPanic)
			test.So(func() {
				golog.Panicln(foo)
			}, test.ShouldPanic)
			test.So(output.log, test.ShouldHaveLength, 5)
		})
	})
}
