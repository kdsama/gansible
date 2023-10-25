package internal

type Engine struct {
	inventory *MainInventory
	playbook  *PlayBook
}

func NewEngine(playbookPath string, hostPath string) *Engine {

	var (
		inventory *MainInventory
		pb        *PlayBook
	)
	if hostPath != "" {
		opts := InventoryPathOptions(hostPath)
		inventory = NewInventory(opts)
	} else {
		inventory = NewInventory()
	}

	pb = NewPlaybook(playbookPath)

	return &Engine{
		inventory: inventory,
		playbook:  pb,
	}
}

func (e *Engine) Run() {
	// os := "ubuntu"
	for i := range e.playbook.Plays {
		e.playbook.Generate(i)
	}
}
