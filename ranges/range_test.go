package ranges

import (
	"math/rand"
	"testing"
	"time"

	"github.com/aikesliu/utils/log"
)

func TestRangeInt64(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	ri := &RangeInt64{
		Max: 10,
	}
	log.SetConsole(true)
	for i := 0; i <= int(ri.Max); i++ {
		ri.Min = int64(i)
		log.D("%v~%v", ri.Min, ri.Max)
		for j := 0; j < 20; j++ {
			log.D("    %v", ri.Rand())
		}
	}
	/*
		=== RUN   TestRangeInt64
		2021/07/19 16:50:01 [debug  ] 0~10
		2021/07/19 16:50:01 [debug  ]     0
		2021/07/19 16:50:01 [debug  ]     1
		2021/07/19 16:50:01 [debug  ]     9
		2021/07/19 16:50:01 [debug  ]     8
		2021/07/19 16:50:01 [debug  ]     0
		2021/07/19 16:50:01 [debug  ]     3
		2021/07/19 16:50:01 [debug  ]     0
		2021/07/19 16:50:01 [debug  ]     0
		2021/07/19 16:50:01 [debug  ]     4
		2021/07/19 16:50:01 [debug  ]     4
		2021/07/19 16:50:01 [debug  ]     1
		2021/07/19 16:50:01 [debug  ]     9
		2021/07/19 16:50:01 [debug  ]     0
		2021/07/19 16:50:01 [debug  ]     6
		2021/07/19 16:50:01 [debug  ]     0
		2021/07/19 16:50:01 [debug  ]     2
		2021/07/19 16:50:01 [debug  ]     4
		2021/07/19 16:50:01 [debug  ]     8
		2021/07/19 16:50:01 [debug  ]     1
		2021/07/19 16:50:01 [debug  ]     7
		2021/07/19 16:50:01 [debug  ] 1~10
		2021/07/19 16:50:01 [debug  ]     4
		2021/07/19 16:50:01 [debug  ]     3
		2021/07/19 16:50:01 [debug  ]     4
		2021/07/19 16:50:01 [debug  ]     2
		2021/07/19 16:50:01 [debug  ]     2
		2021/07/19 16:50:01 [debug  ]     2
		2021/07/19 16:50:01 [debug  ]     9
		2021/07/19 16:50:01 [debug  ]     5
		2021/07/19 16:50:01 [debug  ]     6
		2021/07/19 16:50:01 [debug  ]     3
		2021/07/19 16:50:01 [debug  ]     3
		2021/07/19 16:50:01 [debug  ]     7
		2021/07/19 16:50:01 [debug  ]     2
		2021/07/19 16:50:01 [debug  ]     2
		2021/07/19 16:50:01 [debug  ]     5
		2021/07/19 16:50:01 [debug  ]     2
		2021/07/19 16:50:01 [debug  ]     3
		2021/07/19 16:50:01 [debug  ]     9
		2021/07/19 16:50:01 [debug  ]     4
		2021/07/19 16:50:01 [debug  ]     6
		2021/07/19 16:50:01 [debug  ] 2~10
		2021/07/19 16:50:01 [debug  ]     2
		2021/07/19 16:50:01 [debug  ]     5
		2021/07/19 16:50:01 [debug  ]     4
		2021/07/19 16:50:01 [debug  ]     7
		2021/07/19 16:50:01 [debug  ]     4
		2021/07/19 16:50:01 [debug  ]     7
		2021/07/19 16:50:01 [debug  ]     4
		2021/07/19 16:50:01 [debug  ]     5
		2021/07/19 16:50:01 [debug  ]     3
		2021/07/19 16:50:01 [debug  ]     2
		2021/07/19 16:50:01 [debug  ]     5
		2021/07/19 16:50:01 [debug  ]     2
		2021/07/19 16:50:01 [debug  ]     4
		2021/07/19 16:50:01 [debug  ]     8
		2021/07/19 16:50:01 [debug  ]     5
		2021/07/19 16:50:01 [debug  ]     2
		2021/07/19 16:50:01 [debug  ]     2
		2021/07/19 16:50:01 [debug  ]     3
		2021/07/19 16:50:01 [debug  ]     5
		2021/07/19 16:50:01 [debug  ]     8
		2021/07/19 16:50:01 [debug  ] 3~10
		2021/07/19 16:50:01 [debug  ]     8
		2021/07/19 16:50:01 [debug  ]     3
		2021/07/19 16:50:01 [debug  ]     4
		2021/07/19 16:50:01 [debug  ]     8
		2021/07/19 16:50:01 [debug  ]     7
		2021/07/19 16:50:01 [debug  ]     6
		2021/07/19 16:50:01 [debug  ]     6
		2021/07/19 16:50:01 [debug  ]     4
		2021/07/19 16:50:01 [debug  ]     7
		2021/07/19 16:50:01 [debug  ]     7
		2021/07/19 16:50:01 [debug  ]     5
		2021/07/19 16:50:01 [debug  ]     5
		2021/07/19 16:50:01 [debug  ]     9
		2021/07/19 16:50:01 [debug  ]     3
		2021/07/19 16:50:01 [debug  ]     4
		2021/07/19 16:50:01 [debug  ]     6
		2021/07/19 16:50:01 [debug  ]     4
		2021/07/19 16:50:01 [debug  ]     6
		2021/07/19 16:50:01 [debug  ]     4
		2021/07/19 16:50:01 [debug  ]     6
		2021/07/19 16:50:01 [debug  ] 4~10
		2021/07/19 16:50:01 [debug  ]     4
		2021/07/19 16:50:01 [debug  ]     4
		2021/07/19 16:50:01 [debug  ]     8
		2021/07/19 16:50:01 [debug  ]     6
		2021/07/19 16:50:01 [debug  ]     5
		2021/07/19 16:50:01 [debug  ]     6
		2021/07/19 16:50:01 [debug  ]     7
		2021/07/19 16:50:01 [debug  ]     8
		2021/07/19 16:50:01 [debug  ]     9
		2021/07/19 16:50:01 [debug  ]     4
		2021/07/19 16:50:01 [debug  ]     8
		2021/07/19 16:50:01 [debug  ]     9
		2021/07/19 16:50:01 [debug  ]     9
		2021/07/19 16:50:01 [debug  ]     9
		2021/07/19 16:50:01 [debug  ]     8
		2021/07/19 16:50:01 [debug  ]     9
		2021/07/19 16:50:01 [debug  ]     7
		2021/07/19 16:50:01 [debug  ]     7
		2021/07/19 16:50:01 [debug  ]     8
		2021/07/19 16:50:01 [debug  ]     5
		2021/07/19 16:50:01 [debug  ] 5~10
		2021/07/19 16:50:01 [debug  ]     8
		2021/07/19 16:50:01 [debug  ]     8
		2021/07/19 16:50:01 [debug  ]     6
		2021/07/19 16:50:01 [debug  ]     5
		2021/07/19 16:50:01 [debug  ]     7
		2021/07/19 16:50:01 [debug  ]     5
		2021/07/19 16:50:01 [debug  ]     7
		2021/07/19 16:50:01 [debug  ]     5
		2021/07/19 16:50:01 [debug  ]     8
		2021/07/19 16:50:01 [debug  ]     5
		2021/07/19 16:50:01 [debug  ]     7
		2021/07/19 16:50:01 [debug  ]     5
		2021/07/19 16:50:01 [debug  ]     9
		2021/07/19 16:50:01 [debug  ]     6
		2021/07/19 16:50:01 [debug  ]     5
		2021/07/19 16:50:01 [debug  ]     9
		2021/07/19 16:50:01 [debug  ]     6
		2021/07/19 16:50:01 [debug  ]     5
		2021/07/19 16:50:01 [debug  ]     8
		2021/07/19 16:50:01 [debug  ]     8
		2021/07/19 16:50:01 [debug  ] 6~10
		2021/07/19 16:50:01 [debug  ]     7
		2021/07/19 16:50:01 [debug  ]     7
		2021/07/19 16:50:01 [debug  ]     6
		2021/07/19 16:50:01 [debug  ]     6
		2021/07/19 16:50:01 [debug  ]     8
		2021/07/19 16:50:01 [debug  ]     7
		2021/07/19 16:50:01 [debug  ]     9
		2021/07/19 16:50:01 [debug  ]     9
		2021/07/19 16:50:01 [debug  ]     6
		2021/07/19 16:50:01 [debug  ]     7
		2021/07/19 16:50:01 [debug  ]     6
		2021/07/19 16:50:01 [debug  ]     6
		2021/07/19 16:50:01 [debug  ]     8
		2021/07/19 16:50:01 [debug  ]     8
		2021/07/19 16:50:01 [debug  ]     9
		2021/07/19 16:50:01 [debug  ]     8
		2021/07/19 16:50:01 [debug  ]     8
		2021/07/19 16:50:01 [debug  ]     9
		2021/07/19 16:50:01 [debug  ]     7
		2021/07/19 16:50:01 [debug  ]     9
		2021/07/19 16:50:01 [debug  ] 7~10
		2021/07/19 16:50:01 [debug  ]     9
		2021/07/19 16:50:01 [debug  ]     7
		2021/07/19 16:50:01 [debug  ]     8
		2021/07/19 16:50:01 [debug  ]     7
		2021/07/19 16:50:01 [debug  ]     8
		2021/07/19 16:50:01 [debug  ]     9
		2021/07/19 16:50:01 [debug  ]     9
		2021/07/19 16:50:01 [debug  ]     9
		2021/07/19 16:50:01 [debug  ]     8
		2021/07/19 16:50:01 [debug  ]     7
		2021/07/19 16:50:01 [debug  ]     8
		2021/07/19 16:50:01 [debug  ]     8
		2021/07/19 16:50:01 [debug  ]     9
		2021/07/19 16:50:01 [debug  ]     7
		2021/07/19 16:50:01 [debug  ]     8
		2021/07/19 16:50:01 [debug  ]     8
		2021/07/19 16:50:01 [debug  ]     9
		2021/07/19 16:50:01 [debug  ]     9
		2021/07/19 16:50:01 [debug  ]     9
		2021/07/19 16:50:01 [debug  ]     8
		2021/07/19 16:50:01 [debug  ] 8~10
		2021/07/19 16:50:01 [debug  ]     9
		2021/07/19 16:50:01 [debug  ]     8
		2021/07/19 16:50:01 [debug  ]     8
		2021/07/19 16:50:01 [debug  ]     8
		2021/07/19 16:50:01 [debug  ]     8
		2021/07/19 16:50:01 [debug  ]     9
		2021/07/19 16:50:01 [debug  ]     8
		2021/07/19 16:50:01 [debug  ]     9
		2021/07/19 16:50:01 [debug  ]     9
		2021/07/19 16:50:01 [debug  ]     9
		2021/07/19 16:50:01 [debug  ]     8
		2021/07/19 16:50:01 [debug  ]     8
		2021/07/19 16:50:01 [debug  ]     9
		2021/07/19 16:50:01 [debug  ]     9
		2021/07/19 16:50:01 [debug  ]     9
		2021/07/19 16:50:01 [debug  ]     9
		2021/07/19 16:50:01 [debug  ]     8
		2021/07/19 16:50:01 [debug  ]     8
		2021/07/19 16:50:01 [debug  ]     8
		2021/07/19 16:50:01 [debug  ]     9
		2021/07/19 16:50:01 [debug  ] 9~10
		2021/07/19 16:50:01 [debug  ]     9
		2021/07/19 16:50:01 [debug  ]     9
		2021/07/19 16:50:01 [debug  ]     9
		2021/07/19 16:50:01 [debug  ]     9
		2021/07/19 16:50:01 [debug  ]     9
		2021/07/19 16:50:01 [debug  ]     9
		2021/07/19 16:50:01 [debug  ]     9
		2021/07/19 16:50:01 [debug  ]     9
		2021/07/19 16:50:01 [debug  ]     9
		2021/07/19 16:50:01 [debug  ]     9
		2021/07/19 16:50:01 [debug  ]     9
		2021/07/19 16:50:01 [debug  ]     9
		2021/07/19 16:50:01 [debug  ]     9
		2021/07/19 16:50:01 [debug  ]     9
		2021/07/19 16:50:01 [debug  ]     9
		2021/07/19 16:50:01 [debug  ]     9
		2021/07/19 16:50:01 [debug  ]     9
		2021/07/19 16:50:01 [debug  ]     9
		2021/07/19 16:50:01 [debug  ]     9
		2021/07/19 16:50:01 [debug  ]     9
		2021/07/19 16:50:01 [debug  ] 10~10
		2021/07/19 16:50:01 [debug  ]     10
		2021/07/19 16:50:01 [debug  ]     10
		2021/07/19 16:50:01 [debug  ]     10
		2021/07/19 16:50:01 [debug  ]     10
		2021/07/19 16:50:01 [debug  ]     10
		2021/07/19 16:50:01 [debug  ]     10
		2021/07/19 16:50:01 [debug  ]     10
		2021/07/19 16:50:01 [debug  ]     10
		2021/07/19 16:50:01 [debug  ]     10
		2021/07/19 16:50:01 [debug  ]     10
		2021/07/19 16:50:01 [debug  ]     10
		2021/07/19 16:50:01 [debug  ]     10
		2021/07/19 16:50:01 [debug  ]     10
		2021/07/19 16:50:01 [debug  ]     10
		2021/07/19 16:50:01 [debug  ]     10
		2021/07/19 16:50:01 [debug  ]     10
		2021/07/19 16:50:01 [debug  ]     10
		2021/07/19 16:50:01 [debug  ]     10
		2021/07/19 16:50:01 [debug  ]     10
		2021/07/19 16:50:01 [debug  ]     10
		--- PASS: TestRangeInt64 (0.02s)

	*/
}

func TestRangeU32(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	ri := &RangeUint32{
		Max: 10,
	}
	for i := 0; i <= int(ri.Max); i++ {
		ri.Min = uint32(i)
		log.D("%v~%v", ri.Min, ri.Max)
		for j := 0; j < 20; j++ {
			log.D("    %v", ri.Rand())
		}
	}
	/*
	 */
}

func TestRangeFloat64(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	rf := &RangeFloat64{
		Max: 0.8,
	}
	for min := 0.0; min <= rf.Max; min += 0.1 {
		rf.Min = min
		log.D("%0.1f~%0.1f", rf.Min, rf.Max)
		for i := 0; i < 20; i++ {
			log.D("    %0.3f", rf.Rand())
		}
	}
	/*
		=== RUN   TestRangeFloat64
		2021/07/19 15:47:27 [debug  ] 0.0~0.8
		2021/07/19 15:47:27 [debug  ]     0.344
		2021/07/19 15:47:27 [debug  ]     0.488
		2021/07/19 15:47:27 [debug  ]     0.375
		2021/07/19 15:47:27 [debug  ]     0.070
		2021/07/19 15:47:27 [debug  ]     0.757
		2021/07/19 15:47:27 [debug  ]     0.651
		2021/07/19 15:47:27 [debug  ]     0.684
		2021/07/19 15:47:27 [debug  ]     0.600
		2021/07/19 15:47:27 [debug  ]     0.701
		2021/07/19 15:47:27 [debug  ]     0.495
		2021/07/19 15:47:27 [debug  ]     0.539
		2021/07/19 15:47:27 [debug  ]     0.373
		2021/07/19 15:47:27 [debug  ]     0.243
		2021/07/19 15:47:27 [debug  ]     0.142
		2021/07/19 15:47:27 [debug  ]     0.467
		2021/07/19 15:47:27 [debug  ]     0.596
		2021/07/19 15:47:27 [debug  ]     0.441
		2021/07/19 15:47:27 [debug  ]     0.025
		2021/07/19 15:47:27 [debug  ]     0.488
		2021/07/19 15:47:27 [debug  ]     0.238
		2021/07/19 15:47:27 [debug  ] 0.1~0.8
		2021/07/19 15:47:27 [debug  ]     0.305
		2021/07/19 15:47:27 [debug  ]     0.443
		2021/07/19 15:47:27 [debug  ]     0.662
		2021/07/19 15:47:27 [debug  ]     0.583
		2021/07/19 15:47:27 [debug  ]     0.483
		2021/07/19 15:47:27 [debug  ]     0.274
		2021/07/19 15:47:27 [debug  ]     0.484
		2021/07/19 15:47:27 [debug  ]     0.779
		2021/07/19 15:47:27 [debug  ]     0.397
		2021/07/19 15:47:27 [debug  ]     0.308
		2021/07/19 15:47:27 [debug  ]     0.225
		2021/07/19 15:47:27 [debug  ]     0.596
		2021/07/19 15:47:27 [debug  ]     0.257
		2021/07/19 15:47:27 [debug  ]     0.138
		2021/07/19 15:47:27 [debug  ]     0.164
		2021/07/19 15:47:27 [debug  ]     0.504
		2021/07/19 15:47:27 [debug  ]     0.513
		2021/07/19 15:47:27 [debug  ]     0.172
		2021/07/19 15:47:27 [debug  ]     0.148
		2021/07/19 15:47:27 [debug  ]     0.565
		2021/07/19 15:47:27 [debug  ] 0.2~0.8
		2021/07/19 15:47:27 [debug  ]     0.216
		2021/07/19 15:47:27 [debug  ]     0.578
		2021/07/19 15:47:27 [debug  ]     0.316
		2021/07/19 15:47:27 [debug  ]     0.626
		2021/07/19 15:47:27 [debug  ]     0.210
		2021/07/19 15:47:27 [debug  ]     0.737
		2021/07/19 15:47:27 [debug  ]     0.348
		2021/07/19 15:47:27 [debug  ]     0.614
		2021/07/19 15:47:27 [debug  ]     0.201
		2021/07/19 15:47:27 [debug  ]     0.724
		2021/07/19 15:47:27 [debug  ]     0.503
		2021/07/19 15:47:27 [debug  ]     0.754
		2021/07/19 15:47:27 [debug  ]     0.221
		2021/07/19 15:47:27 [debug  ]     0.395
		2021/07/19 15:47:27 [debug  ]     0.482
		2021/07/19 15:47:27 [debug  ]     0.753
		2021/07/19 15:47:27 [debug  ]     0.454
		2021/07/19 15:47:27 [debug  ]     0.716
		2021/07/19 15:47:27 [debug  ]     0.711
		2021/07/19 15:47:27 [debug  ]     0.584
		2021/07/19 15:47:27 [debug  ] 0.3~0.8
		2021/07/19 15:47:27 [debug  ]     0.591
		2021/07/19 15:47:27 [debug  ]     0.467
		2021/07/19 15:47:27 [debug  ]     0.421
		2021/07/19 15:47:27 [debug  ]     0.515
		2021/07/19 15:47:27 [debug  ]     0.343
		2021/07/19 15:47:27 [debug  ]     0.473
		2021/07/19 15:47:27 [debug  ]     0.330
		2021/07/19 15:47:27 [debug  ]     0.760
		2021/07/19 15:47:27 [debug  ]     0.512
		2021/07/19 15:47:27 [debug  ]     0.561
		2021/07/19 15:47:27 [debug  ]     0.514
		2021/07/19 15:47:27 [debug  ]     0.473
		2021/07/19 15:47:27 [debug  ]     0.700
		2021/07/19 15:47:27 [debug  ]     0.574
		2021/07/19 15:47:27 [debug  ]     0.703
		2021/07/19 15:47:27 [debug  ]     0.351
		2021/07/19 15:47:27 [debug  ]     0.537
		2021/07/19 15:47:27 [debug  ]     0.797
		2021/07/19 15:47:27 [debug  ]     0.571
		2021/07/19 15:47:27 [debug  ]     0.657
		2021/07/19 15:47:27 [debug  ] 0.4~0.8
		2021/07/19 15:47:27 [debug  ]     0.580
		2021/07/19 15:47:27 [debug  ]     0.617
		2021/07/19 15:47:27 [debug  ]     0.767
		2021/07/19 15:47:27 [debug  ]     0.520
		2021/07/19 15:47:27 [debug  ]     0.517
		2021/07/19 15:47:27 [debug  ]     0.734
		2021/07/19 15:47:27 [debug  ]     0.423
		2021/07/19 15:47:27 [debug  ]     0.571
		2021/07/19 15:47:27 [debug  ]     0.459
		2021/07/19 15:47:27 [debug  ]     0.624
		2021/07/19 15:47:27 [debug  ]     0.488
		2021/07/19 15:47:27 [debug  ]     0.761
		2021/07/19 15:47:27 [debug  ]     0.449
		2021/07/19 15:47:27 [debug  ]     0.640
		2021/07/19 15:47:27 [debug  ]     0.567
		2021/07/19 15:47:27 [debug  ]     0.543
		2021/07/19 15:47:27 [debug  ]     0.657
		2021/07/19 15:47:27 [debug  ]     0.562
		2021/07/19 15:47:27 [debug  ]     0.679
		2021/07/19 15:47:27 [debug  ]     0.419
		2021/07/19 15:47:27 [debug  ] 0.5~0.8
		2021/07/19 15:47:27 [debug  ]     0.568
		2021/07/19 15:47:27 [debug  ]     0.785
		2021/07/19 15:47:27 [debug  ]     0.606
		2021/07/19 15:47:27 [debug  ]     0.644
		2021/07/19 15:47:27 [debug  ]     0.547
		2021/07/19 15:47:27 [debug  ]     0.576
		2021/07/19 15:47:27 [debug  ]     0.573
		2021/07/19 15:47:27 [debug  ]     0.552
		2021/07/19 15:47:27 [debug  ]     0.711
		2021/07/19 15:47:27 [debug  ]     0.632
		2021/07/19 15:47:27 [debug  ]     0.579
		2021/07/19 15:47:27 [debug  ]     0.537
		2021/07/19 15:47:27 [debug  ]     0.538
		2021/07/19 15:47:27 [debug  ]     0.775
		2021/07/19 15:47:27 [debug  ]     0.692
		2021/07/19 15:47:27 [debug  ]     0.605
		2021/07/19 15:47:27 [debug  ]     0.706
		2021/07/19 15:47:27 [debug  ]     0.515
		2021/07/19 15:47:27 [debug  ]     0.525
		2021/07/19 15:47:27 [debug  ]     0.544
		2021/07/19 15:47:27 [debug  ] 0.6~0.8
		2021/07/19 15:47:27 [debug  ]     0.606
		2021/07/19 15:47:27 [debug  ]     0.635
		2021/07/19 15:47:27 [debug  ]     0.724
		2021/07/19 15:47:27 [debug  ]     0.717
		2021/07/19 15:47:27 [debug  ]     0.722
		2021/07/19 15:47:27 [debug  ]     0.677
		2021/07/19 15:47:27 [debug  ]     0.758
		2021/07/19 15:47:27 [debug  ]     0.604
		2021/07/19 15:47:27 [debug  ]     0.778
		2021/07/19 15:47:27 [debug  ]     0.636
		2021/07/19 15:47:27 [debug  ]     0.766
		2021/07/19 15:47:27 [debug  ]     0.623
		2021/07/19 15:47:27 [debug  ]     0.716
		2021/07/19 15:47:27 [debug  ]     0.613
		2021/07/19 15:47:27 [debug  ]     0.781
		2021/07/19 15:47:27 [debug  ]     0.708
		2021/07/19 15:47:27 [debug  ]     0.628
		2021/07/19 15:47:27 [debug  ]     0.794
		2021/07/19 15:47:27 [debug  ]     0.670
		2021/07/19 15:47:27 [debug  ]     0.691
		2021/07/19 15:47:27 [debug  ] 0.7~0.8
		2021/07/19 15:47:27 [debug  ]     0.782
		2021/07/19 15:47:27 [debug  ]     0.762
		2021/07/19 15:47:27 [debug  ]     0.796
		2021/07/19 15:47:27 [debug  ]     0.702
		2021/07/19 15:47:27 [debug  ]     0.742
		2021/07/19 15:47:27 [debug  ]     0.724
		2021/07/19 15:47:27 [debug  ]     0.738
		2021/07/19 15:47:27 [debug  ]     0.782
		2021/07/19 15:47:27 [debug  ]     0.790
		2021/07/19 15:47:27 [debug  ]     0.726
		2021/07/19 15:47:27 [debug  ]     0.701
		2021/07/19 15:47:27 [debug  ]     0.739
		2021/07/19 15:47:27 [debug  ]     0.729
		2021/07/19 15:47:27 [debug  ]     0.741
		2021/07/19 15:47:27 [debug  ]     0.742
		2021/07/19 15:47:27 [debug  ]     0.792
		2021/07/19 15:47:27 [debug  ]     0.710
		2021/07/19 15:47:27 [debug  ]     0.734
		2021/07/19 15:47:27 [debug  ]     0.711
		2021/07/19 15:47:27 [debug  ]     0.785
		2021/07/19 15:47:27 [debug  ] 0.8~0.8
		2021/07/19 15:47:27 [debug  ]     0.800
		2021/07/19 15:47:27 [debug  ]     0.800
		2021/07/19 15:47:27 [debug  ]     0.800
		2021/07/19 15:47:27 [debug  ]     0.800
		2021/07/19 15:47:27 [debug  ]     0.800
		2021/07/19 15:47:27 [debug  ]     0.800
		2021/07/19 15:47:27 [debug  ]     0.800
		2021/07/19 15:47:27 [debug  ]     0.800
		2021/07/19 15:47:27 [debug  ]     0.800
		2021/07/19 15:47:27 [debug  ]     0.800
		2021/07/19 15:47:27 [debug  ]     0.800
		2021/07/19 15:47:27 [debug  ]     0.800
		2021/07/19 15:47:27 [debug  ]     0.800
		2021/07/19 15:47:27 [debug  ]     0.800
		2021/07/19 15:47:27 [debug  ]     0.800
		2021/07/19 15:47:27 [debug  ]     0.800
		2021/07/19 15:47:27 [debug  ]     0.800
		2021/07/19 15:47:27 [debug  ]     0.800
		2021/07/19 15:47:27 [debug  ]     0.800
		2021/07/19 15:47:27 [debug  ]     0.800
		--- PASS: TestRangeFloat64 (0.02s)
	*/

}

func TestRangeDuration(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	ri := &RangeDuration{
		Max: 10 * time.Second,
	}
	for i := time.Duration(0); i <= ri.Max; i += time.Second {
		ri.Min = i
		log.D("%v~%v", ri.Min, ri.Max)
		for j := 0; j < 20; j++ {
			log.D("    %v", ri.Rand())
		}
	}
	/*
	 */
}
