package trustlayer.authz

default allow = false

# Разрешаем SuperAdmin любые действия
allow {
    input.user.role == "SuperAdmin"
}

# Разрешаем Auditor только чтение (read-only)
allow {
    input.user.role == "Auditor"
    input.action == "read"
}

# Правило для агентов и остальных: запрещаем опасные операции в нерабочее время (например, DROP/DELETE)
allow {
    input.user.role == "Agent"
    input.action != "drop_database"
    input.action != "force_delete"
}
