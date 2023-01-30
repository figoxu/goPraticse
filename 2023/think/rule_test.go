package think

import (
	"fmt"
	re "github.com/easierway/rule-engine"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = FDescribe("RuleEngine", func() {

	It("Base", func() {
		type Person struct {
			Name  string
			Age   int
			Role  string
			Tasks []string
		}

		fact1 := &Person{
			"Mike", 7, "", []string{},
		}
		rule1 := re.NewRule("Age>=6").
			When(func(f re.Fact) bool {
				p, ok := f.(*Person)
				Ω(ok).Should(BeTrue())
				return p.Age >= 6
			}).
			Then(func(f re.Fact) bool {
				p, ok := f.(*Person)
				Ω(ok).Should(BeTrue())
				if p.Role != "Student" {
					p.Role = "Student"
					return true
				}
				return false
			}).IsRepeatable(true)
		rule2 := re.NewRule("is student").
			When(func(f re.Fact) bool {
				p, ok := f.(*Person)
				Ω(ok).Should(BeTrue())
				return p.Role == "Student"
			}).
			Then(func(f re.Fact) bool {
				p, ok := f.(*Person)
				Ω(ok).Should(BeTrue())
				p.Tasks = append(p.Tasks, "Go to school")
				return true
			})

		rule3 := re.NewRule("School").
			When(func(f re.Fact) bool {
				p, ok := f.(*Person)
				Ω(ok).Should(BeTrue())
				for _, task := range p.Tasks {
					if task == "Go to school" {
						return true
					}
				}
				return false
			}).
			Then(func(f re.Fact) bool {
				p, ok := f.(*Person)
				Ω(ok).Should(BeTrue())
				p.Tasks = append(p.Tasks, "Learning")
				return true
			}).WithOrder(10)
		engine := re.NewEngine()
		engine.AppendRule(rule3).
			AppendRule(rule1).
			AppendRule(rule2).
			Analyze(fact1)
		fmt.Println(*fact1)
		Ω(fact1.Role).Should(BeEquivalentTo("Student"))
		Ω(len(fact1.Tasks) < 2).Should(BeFalse())
	})
})
