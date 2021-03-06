package cfbroker_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/cf-platform-eng/cloudformation-broker/cfbroker"
)

var _ = Describe("Catalog", func() {
	var (
		catalog Catalog

		plan1 = ServicePlan{ID: "Plan-1"}
		plan2 = ServicePlan{ID: "Plan-2"}

		service1 = Service{ID: "Service-1", Plans: []ServicePlan{plan1}}
		service2 = Service{ID: "Service-2", Plans: []ServicePlan{plan2}}
	)

	Describe("Validate", func() {
		BeforeEach(func() {
			catalog = Catalog{}
		})

		It("does not return error if all fields are valid", func() {
			err := catalog.Validate()

			Expect(err).ToNot(HaveOccurred())
		})

		It("returns error if Services are not valid", func() {
			catalog.Services = []Service{
				Service{},
			}

			err := catalog.Validate()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Validating Services configuration"))
		})
	})

	Describe("FindService", func() {
		BeforeEach(func() {
			catalog = Catalog{
				Services: []Service{service1, service2},
			}
		})

		It("returns true and the Service if it is found", func() {
			service, found := catalog.FindService("Service-1")
			Expect(service).To(Equal(service1))
			Expect(found).To(BeTrue())
		})

		It("returns false if it is not found", func() {
			_, found := catalog.FindService("Service-?")
			Expect(found).To(BeFalse())
		})
	})

	Describe("FindServicePlan", func() {
		BeforeEach(func() {
			catalog = Catalog{
				Services: []Service{service1, service2},
			}
		})

		It("returns true and the Service Plan if it is found", func() {
			plan, found := catalog.FindServicePlan("Plan-1")
			Expect(plan).To(Equal(plan1))
			Expect(found).To(BeTrue())
		})

		It("returns false if it is not found", func() {
			_, found := catalog.FindServicePlan("Plan-?")
			Expect(found).To(BeFalse())
		})
	})
})

var _ = Describe("Service", func() {
	var (
		service Service

		validService = Service{
			ID:              "Service-1",
			Name:            "Service 1",
			Description:     "Service 1 description",
			Bindable:        true,
			Tags:            []string{"service"},
			Metadata:        &ServiceMetadata{},
			Requires:        []string{"syslog"},
			PlanUpdateable:  true,
			Plans:           []ServicePlan{},
			DashboardClient: &DashboardClient{},
		}
	)

	BeforeEach(func() {
		service = validService
	})

	Describe("Validate", func() {
		It("does not return error if all fields are valid", func() {
			err := service.Validate()
			Expect(err).ToNot(HaveOccurred())
		})

		It("returns error if ID is empty", func() {
			service.ID = ""

			err := service.Validate()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Must provide a non-empty ID"))
		})

		It("returns error if Name is empty", func() {
			service.Name = ""

			err := service.Validate()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Must provide a non-empty Name"))
		})

		It("returns error if Description is empty", func() {
			service.Description = ""

			err := service.Validate()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Must provide a non-empty Description"))
		})

		It("returns error if Plans are not valid", func() {
			service.Plans = []ServicePlan{
				ServicePlan{},
			}

			err := service.Validate()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Validating Plans configuration"))
		})
	})
})

var _ = Describe("ServicePlan", func() {
	var (
		servicePlan ServicePlan

		validServicePlan = ServicePlan{
			ID:          "Plan-1",
			Name:        "Plan 1",
			Description: "Plan-1 description",
			Metadata:    &ServicePlanMetadata{},
			Free:        true,
			CloudFormationProperties: CloudFormationProperties{
				TemplateURL: "test-template-url",
			},
		}
	)

	BeforeEach(func() {
		servicePlan = validServicePlan
	})

	Describe("Validate", func() {
		It("does not return error if all fields are valid", func() {
			err := servicePlan.Validate()
			Expect(err).ToNot(HaveOccurred())
		})

		It("returns error if ID is empty", func() {
			servicePlan.ID = ""

			err := servicePlan.Validate()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Must provide a non-empty ID"))
		})

		It("returns error if Name is empty", func() {
			servicePlan.Name = ""

			err := servicePlan.Validate()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Must provide a non-empty Name"))
		})

		It("returns error if Description is empty", func() {
			servicePlan.Description = ""

			err := servicePlan.Validate()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Must provide a non-empty Description"))
		})
	})
})

var _ = Describe("CloudFormationProperties", func() {
	var (
		cloudformationProperties CloudFormationProperties

		validCloudFormationProperties = CloudFormationProperties{
			TemplateURL: "test-template-url",
		}
	)

	BeforeEach(func() {
		cloudformationProperties = validCloudFormationProperties
	})

	Describe("Validate", func() {
		It("does not return error if all fields are valid", func() {
			err := cloudformationProperties.Validate()
			Expect(err).ToNot(HaveOccurred())
		})

		It("returns error if OnFailure is empty", func() {
			cloudformationProperties.OnFailure = "unknown"

			err := cloudformationProperties.Validate()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("OnFailure 'unknown' not supported"))
		})

		It("returns error if TemplateURL is empty", func() {
			cloudformationProperties.TemplateURL = ""

			err := cloudformationProperties.Validate()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Must provide a non-empty TemplateURL"))
		})
	})
})
