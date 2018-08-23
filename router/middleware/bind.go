package middleware

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-macaron/binding"
	"github.com/go-macaron/session"
	"github.com/kindlyfire/golog/modules/context"
	macaron "gopkg.in/macaron.v1"
)

func Bind(obj interface{}) macaron.Handler {
	handler := binding.BindIgnErr(obj)

	return func(ctx *context.Context) {
		ctx.RealCtx.Invoke(handler)
		ctx.RealCtx.Invoke(func(ctx *context.Context, errs binding.Errors, flash *session.Flash, sess session.Store) {
			if len(errs) == 0 {
				return
			}

			// We only send a flash message for the first error
			err := errs[0]
			t := reflect.TypeOf(obj)

			if err.Classification == "Custom" {
				flash.Error(err.Message)
			} else {
				field, ok := t.FieldByName(err.FieldNames[0])

				if !ok {
					return
				}

				errMsg := field.Tag.Get("err")

				if errMsg != "" {
					flash.Error(errMsg)
				}
			}

			value := ctx.RealCtx.Injector.GetVal(t)
			forwardKeys := []string{}

			// Forward all form fields that have forwardAs:"" non-empty
			for i := 0; i < t.NumField(); i++ {
				field := t.Field(i)
				tag := field.Tag.Get("forwardAs")

				if tag != "" {
					sess.Set(tag, value.Field(i).String())
					forwardKeys = append(forwardKeys, tag)
				}
			}

			sess.Set("Forwards", strings.Join(forwardKeys, ","))

			// Redirect back from where we came from
			ctx.Redirect(ctx.RealCtx.Req.URL.Path)
		})
	}
}

func BindErrForwarder(ctx *context.Context, sess session.Store) {
	forwardsStr, ok := sess.Get("Forwards").(string)

	// There is nothing to forward
	if !ok {
		return
	}

	forwards := strings.Split(forwardsStr, ",")

	for _, fw := range forwards {
		str, ok := sess.Get(fw).(string)

		if !ok {
			continue
		}

		ctx.Data[fw] = str
	}

	fmt.Println(forwards)
}
