package providers

func List() []any {
	return []any{
		NewConfigServiceProvider(),
		NewCliServiceProvider(),
		NewDatabaseServiceProvider(),
		NewRepositoryProvider(),
		NewServiceProvider(),
		NewUseCaseProvider(),
		NewControllereProvider(),
	}
}
