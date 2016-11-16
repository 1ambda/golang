package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

func main() {
	/**
	 * If you write a function that takes an `interface{}` value as a param,
	 * you can supply that function with any value.
	 *
	 * `func DoSomething(v interface{})`
	 *
	 * `v` is not any type. It is of `interface{}` type.
	 * When passing a value into `DoSomething` function,
	 * the Go runtime will perform a type conversino (if necessary),
	 * and convert the value to an interface{} value.
	 *
	 *
	 * All values have exactly one type at runtime, and v's static type is `interface{}`
	 *
	 * An interface value (passed to as a param) have two words of data.
	 * - one word is used to point to a method table for the underlying type
	 * - the other word is used to point the actual data being held by that value
	 *
	 * We can't convert `T[]` to `[]interface{}` for this reason.
	 * They have different memory representations.
	 */

	names := []string{"stanley", "david", "oscar"}
	vals := make([]interface{}, len(names))
	for i, v := range names {
		vals[i] = v
	}
	_ = vals

	/**
	 * Pointer type can access to methods of value type while
	 * value type can't access to methods of pointer type
	 *
	 * Usually, people don't use value receiver since
	 * everyting passed as a param will be copied.
	 */

	animals := []Animal{&Dog{}, new(Cat), Llama{}, JavaProgrammer{}}

	for _, animal := range animals {
		_ = animal
	}

	// json parsing
	var val map[string]interface{}
	input := `
{
	"created_at": "Thu May 31 00:00:01 +0000 2012"
}
`
	if err := json.Unmarshal([]byte(input), &val); err != nil {
		panic(err)
	}

	fmt.Println(val)
	for k, v := range val {
		fmt.Println(k, v, reflect.TypeOf(v))
	}

	var val2 map[string]Timestamp

	if err := json.Unmarshal([]byte(input), &val2); err != nil {
		panic(err)
	}

	fmt.Println(val2)
	for k, v := range val2 {
		fmt.Println(k, reflect.TypeOf(v))
	}
	fmt.Println(time.Time(val2["created_at"]))
}

type Animal interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct{}

func (c *Cat) Speak() string {
	return "Meow!"
}

type Llama struct{}

func (l Llama) Speak() string {
	return "????"
}

type JavaProgrammer struct{}

func (j JavaProgrammer) Speak() string {
	return "Design patterns!"
}

type Timestamp time.Time

func (t *Timestamp) UnmarshalJSON(b []byte) error {
	v, err := time.Parse(time.RubyDate, string(b[1:len(b)-1]))
	if err != nil {
		return err
	}
	*t = Timestamp(v)
	return nil
}
