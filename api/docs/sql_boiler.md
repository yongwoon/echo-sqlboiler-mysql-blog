# Sql Boiler

## Init sql boiler

1. `create sqlboiler.toml` file on root directory
2. execute command as blow

```bash
# Format
sqlboiler mysql --wipe
sqlboiler mysql --output [OUTPUT PATH] --pkgname [PACKAGE NAME] --wipe
```

```bash
# In this project
sqlboiler mysql --output "db/orm" --pkgname orm --wipe
```

## Sql boiler を叩く時

- 最初の sql boiler 導入時(model があったら)
- model 追加後
