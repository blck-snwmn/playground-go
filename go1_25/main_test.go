package main

import (
	jsonv1 "encoding/json"
	jsonv2 "encoding/json/v2"
	"fmt"
	"sync"
	"testing"
	"testing/synctest"
	"time"
)

func TestSynctest(t *testing.T) {
	w := t.Output()
	enc := jsonv1.NewEncoder(w)
	enc.SetIndent("", "  ")
	enc.Encode(map[string]string{
		"test": "synctest",
	})

	t.Attr("test", "synctest")

	// This test is set to wait for one day using sleep, but the test execution time is less than one second.
	synctest.Test(t, func(t *testing.T) {
		start := time.Now()
		time.Sleep(24 * time.Hour)

		since := time.Since(start)
		if since != 24*time.Hour {
			t.Errorf("time.Since(start) = %v, want >= 1s", since)
		}
	})
}

func TestJSONV2Deterministic(t *testing.T) {

	m := map[string]string{}
	for i := 0; i < 100; i++ {
		m[fmt.Sprintf("key%d", i)] = fmt.Sprintf("value%d", i)
	}

	var wg sync.WaitGroup
	jsonbs := make([]string, 100)
	for i := 0; i < 100; i++ {
		wg.Go(func() {
			b, err := jsonv2.Marshal(m, jsonv2.Deterministic(true))
			if err != nil {
				t.Error("Error marshaling JSON:", err)
				return
			}
			jsonbs[i] = string(b)
		})
	}
	wg.Wait()

	for _, b := range jsonbs[1:] {
		if b != jsonbs[0] {
			t.Error("JSON output is not deterministic")
			return
		}
	}
	t.Logf("JSON output is deterministic: %s", jsonbs[0])
	// example output: JSON output is deterministic: {"key0":"value0","key1":"value1","key10":"value10","key11":"value11","key12":"value12","key13":"value13","key14":"value14","key15":"value15","key16":"value16","key17":"value17","key18":"value18","key19":"value19","key2":"value2","key20":"value20","key21":"value21","key22":"value22","key23":"value23","key24":"value24","key25":"value25","key26":"value26","key27":"value27","key28":"value28","key29":"value29","key3":"value3","key30":"value30","key31":"value31","key32":"value32","key33":"value33","key34":"value34","key35":"value35","key36":"value36","key37":"value37","key38":"value38","key39":"value39","key4":"value4","key40":"value40","key41":"value41","key42":"value42","key43":"value43","key44":"value44","key45":"value45","key46":"value46","key47":"value47","key48":"value48","key49":"value49","key5":"value5","key50":"value50","key51":"value51","key52":"value52","key53":"value53","key54":"value54","key55":"value55","key56":"value56","key57":"value57","key58":"value58","key59":"value59","key6":"value6","key60":"value60","key61":"value61","key62":"value62","key63":"value63","key64":"value64","key65":"value65","key66":"value66","key67":"value67","key68":"value68","key69":"value69","key7":"value7","key70":"value70","key71":"value71","key72":"value72","key73":"value73","key74":"value74","key75":"value75","key76":"value76","key77":"value77","key78":"value78","key79":"value79","key8":"value8","key80":"value80","key81":"value81","key82":"value82","key83":"value83","key84":"value84","key85":"value85","key86":"value86","key87":"value87","key88":"value88","key89":"value89","key9":"value9","key90":"value90","key91":"value91","key92":"value92","key93":"value93","key94":"value94","key95":"value95","key96":"value96","key97":"value97","key98":"value98","key99":"value99"}
}
