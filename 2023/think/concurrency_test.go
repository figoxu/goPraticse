package think_test

import (
	"context"
	concurrency "github.com/easierway/concurrency_utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"
)

var _ = Describe("concurrency task like java future", func() {

	It("Pure", func() {
		mockResult := [5]int{1, 2, 3, 4, 5}
		mockTimeOutFn := func(ctx context.Context) concurrency.TaskResult {
			time.Sleep(time.Second * 1)
			return concurrency.TaskResult{
				Result: &mockResult,
			}
		}
		start := time.Now()
		retStub := concurrency.AsynExecutor(context.TODO(), mockTimeOutFn, 2000)
		Ω(time.Since(start) < time.Second).Should(BeTrue()) // 类似Java Future的效果
		Ω(retStub.GetResult().Err).Should(BeNil())
		cost := time.Since(start)
		Ω(cost < time.Second*time.Duration(2)).Should(BeTrue())
		Ω(cost > time.Second).Should(BeTrue())
		ret, ok := retStub.GetResult().Result.(*[5]int)
		Ω(ok).Should(BeTrue())
		Ω(ret[4]).Should(Equal(mockResult[4]))
	})
})
