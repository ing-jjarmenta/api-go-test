# api-go-test

## 🗄️ MongoDB Adapter en Go

Este proyecto implementa una capa de **abstracción y adaptación** sobre el driver oficial de MongoDB (`mongo-driver/v2`) para:

- Desacoplar la lógica de negocio del driver.
- Permitir **pruebas unitarias** sin necesidad de un Mongo real.
- Mantener un diseño limpio con **interfaces, adapters y mocks**.

---

## 📦 Estructura principal

/mongodb
├── mongo.go # Interfaces + wrappers (connectFunc, pingFunc)
├── mongo_adapter.go # Adapters que envuelven al driver real
├── mongo_mocks_test.go # Mocks con testify
├── mongo_test.go # Unit tests de NewMongoClient y TasksCollection
├── mongo_wrappers_test.go # Tests dummy para cubrir wrappers
└──

## ⚡ Patrón usado: Adapter + Interfaces

- `MongoClient`, `MongoDatabase`, `MongoCollection` → **interfaces**.
- `AdapterClient`, `AdapterDatabase`, `AdapterCollection` → **wrappers del driver real**.
- `MockClient`, `MockDatabase`, `MockCollection` → **mocks de prueba con testify**.
