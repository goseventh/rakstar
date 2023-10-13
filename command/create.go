package command

// Command seta o nome do comando para ser invocado pelo Player.
// O nome não deve conter espaços ou caracteres acentuados ou 
// especiais 
func (cb *commandBuilder) Command(cmd string) *commandBuilder {
	cb.cmd = cmd
	return cb
}

// Alias seta um sinônimo para o nome principal do comando
func (cb *commandBuilder) Alias(cmd string) *commandBuilder {
	cb.alias = append(cb.alias, cmd)
	return cb
}

// Handler seta a função que manipulará o comando, estabelecendo
// lógicas e regras de negócio 
func (cb *commandBuilder) Handler(handler CommandHandler) *commandBuilder {
	cb.handler = handler
	return cb
}

// Create registra o comando na lista de comandos, 
// para se tornar válido
func (cb *commandBuilder) Create() {
	cmd := &Command{
		Name:       cb.cmd,
		Aliases:    cb.alias,
		Handler:    cb.handler,
		conditions: cb.conditionals,
	}

	RegisterCommand(cmd)
}
