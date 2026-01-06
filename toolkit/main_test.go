package toolkit

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/samber/do/v2"
	"testing"
	"time"
)

func TestNil(t *testing.T) {
	type Boy struct {
		Name *string `json:"name"`
		Age  int     `json:"age"`
	}

	jsonString := `{"age":18}`
	boy := &Boy{}
	err := json.Unmarshal([]byte(jsonString), boy)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(boy)

}

func TestNil2(t *testing.T) {
	type Boy struct {
		Name *string `json:"name"`
		Age  int     `json:"age"`
	}

	boy := &Boy{
		Name: nil,
		Age:  20,
	}

	bytes, err := json.Marshal(boy)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bytes))

}

func TestGoro(t *testing.T) {
	var err error
	go func() {
		err = errors.New("协程中错误:测试错误")
	}()

	time.Sleep(1 * time.Second)
	if err != nil {
		fmt.Println(err)
		return
	}
}

type BirdIF interface {
	say()
}

type PersonIF interface {
	say()
}

type Boy struct {
	name string
}

func (r Boy) say() {
	fmt.Println("boy")
}

type Girl struct {
	name string
}

func (r Girl) say() {
	fmt.Println("girl")
}

func TestDo(t *testing.T) {
	injector := do.New()
	//do.Provide(injector, func(injector do.Injector) (PersonIF, error) {
	//	return &Boy{}, nil
	//})
	//do.Provide(injector, func(injector do.Injector) (PersonIF, error) {
	//	return &Boy{}, nil
	//})
	do.ProvideNamed(injector, "girl", func(injector do.Injector) (PersonIF, error) {
		return &Girl{
			name: "lily",
		}, nil
	})

	//invoke, err := do.Invoke[PersonIF](injector)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	invoke, err := do.InvokeNamed[BirdIF](injector, "girl")

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(invoke)
}
