package composite

// IOrganization organization interface
type IOrganization interface {
	Count() int
}

// Employee employee struct
type Employee struct {
	Name string
}

// Count employee function
func (Employee) Count() int {
	return 1
}

// Departement departement struct
type Departement struct {
	Name             string
	SubOrganizations []IOrganization
}

// Count count departement employee
func (d Departement) Count() int {
	c := 0
	for _, org := range d.SubOrganizations {
		c += org.Count()
	}
	return c
}

// AddSub add sub organization
func (d *Departement) AddSub(org IOrganization) {
	d.SubOrganizations = append(d.SubOrganizations, org)
}

// NewOrganization create an new organization
func NewOrganization() IOrganization {
	root := &Departement{Name: "root"}
	for i := 0; i < 10; i++ {
		root.AddSub(&Employee{})
		root.AddSub(&Departement{Name: "sub", SubOrganizations: []IOrganization{&Employee{}}})
	}
	return root
}
