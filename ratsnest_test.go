package ratsnest

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"errors"
	. "github.com/onsi/ginkgo/extensions/table"
	"strings"
)

var _ = Describe("Rat's Nest", func() {
	var (
		n *Node
		data map[string]interface{}
		root *Node
		err error
	)

	BeforeEach(func() {
		n = nil
		data = map[string]interface{}{
			"lorem": "ipsum",
			"dolor": 592,
			"amet": map[string]interface{}{
				"consectetur": "adipiscing",
				"elit": []interface{}{
					"donec",
					"hendrerit",
					"turpis",
					map[string]interface{}{
						"vel": "sem",
					},
				},
				"gravida": true,
				"vestibulum": []interface{}{
					36954.02,
					true,
					72,
				},
			},
		}
		err = nil
	})

	Specify("initialization with a nil map results in an error", func() {
		_, err = New(nil)
		Expect(err).To(MatchError(errors.New("No source data was found. Be sure to initialize ratnest with your data by using ratsnest.New(...)")))
	})

	Describe("Node validation", func() {
		BeforeEach(func() {
			n = &Node{
				Value: "foobar",
				Key: "bazbat",
				sourceData: map[string]interface{}{
					"foo": "bar",
				},
			}
		})

		JustBeforeEach(func() {
			err = n.isValid()
		})

		Context("without a key nor value", func() {
			BeforeEach(func() {
				n.Key = ""
				n.Value = nil
			})

			It("returns an error", func() {
				Expect(err).To(MatchError(errors.New("Nodes must have a key or a value, or both")))
			})
		})

		Context("with a negative MaxDepth", func() {
			BeforeEach(func() {
				n.MaxDepth = -1
			})

			It("returns an error", func() {
				Expect(err).To(MatchError(errors.New("MaxDepth must be a positive integer")))
			})
		})

		Context("without sourceData", func() {
			BeforeEach(func() {
				n.sourceData = nil
			})

			It("returns an error", func() {
				Expect(err).To(MatchError(errors.New("No source data was found. Be sure to initialize ratnest with your data by using ratsnest.New(...)")))
			})
		})
	})

	Describe("registering new child node requirements", func() {
		var reqNode *Node

		BeforeEach(func() {
			n = &Node{
				Value: "ipsum",
				Key: "lorem",
			}
			root, err = New(data)
			Expect(err).NotTo(HaveOccurred())
		})

		JustBeforeEach(func() {
			reqNode, err = root.Require(*n)
		})

		Context("when the required node is invalid", func() {
			BeforeEach(func() {
				n.Value = nil
				n.Key = ""
			})

			Specify("an error is returned", func() {
				Expect(err).To(MatchError(errors.New("Nodes must have a key or a value, or both")))
			})
		})

		Context("when all is right in the world", func() {
			Specify("happy times abound", func() {
				Expect(err).NotTo(HaveOccurred())
			})

			Specify("the root node has a child", func() {
				Expect(len(root.childRequirements)).To(Equal(1))
			})
		})

		Describe("validation of data presence", func() {
			Context("with valid requirements", func() {
				type (
					validationCase struct {
						toRegister Node
						expErr error
						noChildRequirements bool
					}
					
					validationCases []validationCase
				)

				satMatcher := func(cases validationCases) {
					pn := root
					for _, c := range cases {
						rn, err := pn.Require(c.toRegister)
						if c.expErr == nil {
							Expect(err).NotTo(HaveOccurred())
						} else {
							Expect(err).To(MatchError(c.expErr))
						}
						c.toRegister.Case = CaseInsensitive
						c.toRegister.Key = strings.ToUpper(c.toRegister.Key)
						valStr, valIsStr := c.toRegister.Value.(string)
						if valIsStr {
							c.toRegister.Value = strings.ToUpper(valStr)
						}
						rn, err = pn.Require(c.toRegister)
						if c.expErr == nil {
							Expect(err).NotTo(HaveOccurred())
						} else {
							Expect(err).To(MatchError(c.expErr))
						}
						if !c.noChildRequirements {
							pn = rn
						}
					}
				}

				DescribeTable("satisfaction validation", satMatcher,
					Entry("for an existing requirement with a depth of 1", validationCases{
						{
							toRegister: Node{Value: 592},
						},
					}),
					Entry("for an existing requirement with a depth of 2", validationCases{
						{
							toRegister: Node{Value: "adipiscing"},
						},
					}),
					Entry("for an existing requirement with a depth of 3", validationCases{
						{
							toRegister: Node{Value: "donec"},
						},
					}),
					Entry("for an existing requirement with a depth of 4", validationCases{
						{
							toRegister: Node{Value: "sem"},
						},
					}),
					Entry("for an existing requirement with a depth of 1 and a key", validationCases{
						{
							toRegister: Node{
								Value: 592,
								Key: "dolor",
							},
						},
					}),
					Entry("for an existing requirement with a depth of 2 and a key", validationCases{
						{
							toRegister: Node{
								Value: "adipiscing",
								Key: "consectetur",
							},
						},
					}),
					Entry("for an existing requirement with a depth of 3 and a key", validationCases{
						{
							toRegister: Node{
								Value: "donec",
								Key: "elit",
							},
						},
					}),
					Entry("for an existing requirement with a depth of 4 and a key", validationCases{
						{
							toRegister: Node{
								Value: "sem",
								Key: "vel",
							},
						},
					}),
					Entry("for an existing requirement with a depth of 1 and only a key", validationCases{
						{
							toRegister: Node{
								Key: "dolor",
							},
						},
					}),
					Entry("for an existing requirement with a depth of 2 and only a key", validationCases{
						{
							toRegister: Node{
								Key: "consectetur",
							},
						},
					}),
					Entry("for an existing requirement with a depth of 3 and only a key", validationCases{
						{
							toRegister: Node{
								Key: "elit",
							},
						},
					}),
					Entry("for an existing requirement with a depth of 4 and only a key", validationCases{
						{
							toRegister: Node{
								Key: "vel",
							},
						},
					}),
					Entry("for two existing requirements with depths of 1", validationCases{
						{
							toRegister: Node{
								Value: 592,
							},
							noChildRequirements: true,
						},
						{
							toRegister: Node{
								Value: "adipiscing",
							},
						},
					}),
					Entry("for three existing requirements with depths of 1", validationCases{
						{
							toRegister: Node{
								Key: "amet",
							},
						},
						{
							toRegister: Node{
								Key: "elit",
							},
						},
						{
							toRegister: Node{
								Value: "donec",
							},
						},
					}),
					Entry("for four existing requirements with depths of 1", validationCases{
						{
							toRegister: Node{Key: "amet"},
						},
						{
							toRegister: Node{Key: "elit"},
						},
						{
							toRegister: Node{
								Value: map[string]interface{}{
									"vel": "sem",
								},
							},
						},
						{
							toRegister: Node{Value: "sem"},
						},
					}),
					Entry("for requesting an array value", validationCases{
							{
								toRegister: Node{
									Value: []interface{}{
										true,
										72,
										36954.02,
									},
								},
							},
					}),
					Entry("for requesting an array with a mismatched value", validationCases{
							{
								toRegister: Node{
									Value: []interface{}{
										true,
										72,
										"foobar",
									},
								},
								expErr: NodeNotFoundError{},
							},
					}),
					Entry("when requesting a map value with a mismatched key", validationCases{
						{
							toRegister: Node{
								Value: map[string]interface{}{
									"vel": "bar",
								},
							},
							expErr: NodeNotFoundError{},
						},
					}),
					Entry("for two existing requirements with depths over 1", 	validationCases{
							{
								toRegister: Node{Key: "amet"},
							},
							{
								toRegister: Node{Value: "hendrerit"},
							},
					}),
					Entry("for two existing requirements with depths over 1 where one has a maximum depth too low to be satisfied", validationCases{
						{
							toRegister: Node{Key: "amet"},
						},
						{
							toRegister: Node{
								Value: "hendrerit",
								MaxDepth: 1,
							},
							expErr: NodeNotFoundError{},
						},
					}),
					Entry("for two existing requirements with depths over 1 with only keys", validationCases{
						{
							toRegister: Node{Key: "amet"},
						},
						{
							toRegister: Node{Key: "vel"},
						},
					}),
					Entry("for non-existent requirements", validationCases{
						{
							toRegister: Node{Key: "amet"},
						},
						{
							toRegister: Node{Key: "foobar"},
							expErr: NodeNotFoundError{},
						},
					}),
				)
			})
		})
	})

	Describe("error descriptions", func() {
		It("will return the correct error for no criteria", func() {
			Expect(NoCriteriaError{}.Error()).To(Equal("No requirements have been added to the root node."))
		})
		It("will return the correct error for unfound nodes", func() {
			Expect(NodeNotFoundError{}.Error()).To(Equal("The required node was not found in the parent."))
		})
	})
})
