	fails    [][]byte
			fails:    [][]byte{[]byte("notcorrect")},
			fails:    [][]byte{[]byte("notcorrect")},
			fails:    [][]byte{[]byte("notcorrect"), []byte("--- s"), []byte("short")},
			fails:    [][]byte{[]byte("notcorrect"), []byte("+++ s"), []byte("short")},
			fails:    [][]byte{[]byte("notcorrect")},
			fails:    [][]byte{[]byte("notcorrect")},
			fails:    [][]byte{[]byte("notcorrect")},
			fails:    [][]byte{[]byte("notcorrect")},
		for _, fail := range test.fails {
			if test.function(fail) {
				t.Errorf("%s: Parser did not recognize incorrect line.", name)
			}