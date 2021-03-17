package objectid

import (
	"testing"
)

func Test_Generate_One(t *testing.T) {
	id := New().Generate()
	if len(id) != 12 {
		t.Error(id, len(id))
	}
}

func Test_Generate_Multi(t *testing.T) {
	// generator := New()
	// var ids [12]string
	// for i := 0; i < len(ids); i++ {
	// 	ids[i] = generator.Generate()
	// }

	// res := underscore.Chain(ids).Group(func(r string, _ int) string {
	// 	return r
	// }).Size().Value()
	// if res.(int) != len(ids) {
	// 	t.Error(ids)
	// }
}

func Test_Generate_Multi_Async(t *testing.T) {
	// generator := New()
	// var ids [12]string
	// var wg sync.WaitGroup
	// wg.Add(len(ids))
	// for i := 0; i < len(ids); i++ {
	// 	go func(index int) {
	// 		ids[index] = generator.Generate()
	// 		wg.Done()
	// 	}(i)
	// }

	// wg.Wait()

	// res := underscore.Chain(ids).Group(func(r string, _ int) string {
	// 	return r
	// }).Size().Value()
	// if res.(int) != len(ids) {
	// 	t.Error(ids)
	// }
}
