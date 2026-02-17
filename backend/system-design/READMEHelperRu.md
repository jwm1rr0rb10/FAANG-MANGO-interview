## Масштабируемость 
- это свойство системы справляться с растущим объемом работы. Одно из определений программных систем указывает, что это может быть достигнуто путем добавления ресурсов в систему.

## Надежность 
- это последовательность, достоверность и воспроизводимость измерений, систем или продуктов во времени. Она гарантирует воспроизводимость результатов в идентичных условиях (статистика) или то, что компонент выполняет свою функцию без сбоев в течение заданного периода времени (инженерные расчеты).

## Компромисс
- баланс, достигаемый между двумя желательными, но несовместимыми характеристиками; компромисс.

## Узкое место 
- это точка перегрузки в производственной системе (например, на сборочной линии или в компьютерной сети), которая останавливает или значительно замедляет работу системы.

# 📚 System Design & Distributed Systems Knowledge Base

Полный конспект по системному дизайну, базам данных, распределённым системам и отказоустойчивости.

---

# 1️⃣ Свойства информационных систем

### Надёжность
Способность системы функционировать без сбоев в течение требуемого времени.  
Включает:
- Отказоустойчивость
- Восстановление после ошибок
- Гарантию корректности данных

### Масштабируемость
Способность увеличивать производительность при добавлении ресурсов:
- Вертикальное (scale-up)
- Горизонтальное (scale-out)

### Производительность
- Время отклика (latency)
- Пропускная способность (throughput)
- Эффективность использования CPU / RAM / Disk

### Удобство сопровождения
- Модульность
- Документированность
- Тестируемость
- Простота изменений

### Безопасность
- Аутентификация
- Авторизация
- Шифрование
- Аудит
- Защита от атак и утечек

---

# 2️⃣ Критерии информационных систем

### Data / Compute Intensive
- **Data-intensive** — узкое место в объёме данных (аналитические БД)
- **Compute-intensive** — узкое место в вычислениях (научные расчёты)

### Read / Write Intensive
- Read-heavy (новостные порталы)
- Write-heavy (логирование)

### Low Latency
- Миллисекундные задержки
- Онлайн-игры
- Финансовые транзакции

### High Throughput
- Миллионы запросов в секунду
- Потоковая обработка событий

---

# 3️⃣ Балансировка нагрузки

## Виды
- Клиентская
- Серверная

## DNS / GeoDNS
- Геораспределение
- Медленное обновление

## L4 / L7
- **L4** — TCP/UDP
- **L7** — HTTP/HTTPS (анализ содержимого)

## Алгоритмы
- Round Robin
- Least Connections
- IP Hash
- Random
- Weighted

---

# 4️⃣ Проксирование

## Forward / Reverse Proxy
- Forward — для клиентов
- Reverse — перед серверами (например, Nginx)

## Требования
- Функциональные
- Нефункциональные (SLA, безопасность, производительность)

## Расчёт нагрузки
- RPS
- Пиковая нагрузка
- Время обработки

---

# 5️⃣ Кэширование

## Типы
- Внутренний
- Внешний (Redis, Memcached)

## Метрики
- Hit ratio
- Miss ratio

## Thundering Herd Problem
Решения:
- Блокировки
- Random TTL
- Background refresh

## Алгоритмы вытеснения

- OPT (Belady)
- Random
- LRU
- SLRU
- TLRU
- LRU-k
- MRU
- LFU
- FIFO / LIFO
- 2Q
- Second Chance
- Clock

## Инвалидация
- TTL
- По событию
- Версионирование
- Тегирование

## Паттерны
- Cache Aside
- Cache Through
- Cache Ahead

---

# 6️⃣ API

## CRUD
Create / Read / Update / Delete

## Under / Over Fetching
- Under — не хватает данных
- Over — лишние данные

## Подходы
- SOAP
- REST
- RPC (gRPC)
- GraphQL

## Обновления
- Polling
- Long Polling
- SSE
- WebSockets

## Версионирование
- URL
- Заголовки
- Параметры

## Идемпотентность
Повторный запрос не меняет состояние.

---

# 7️⃣ Observability

### Три столпа
- Метрики
- Логи
- Трейсы

### Алертинг
Оповещения при превышении порогов.

### Continuous Profiling
Постоянный сбор данных о CPU / памяти.

---

# 8️⃣ Виды баз данных

- Реляционные (PostgreSQL, MySQL)
- Документоориентированные (MongoDB)
- Key-Value (Redis, DynamoDB)
- Time-series (InfluxDB)
- Колоночные (ClickHouse)
- Wide-column (Cassandra, HBase)
- Object Storage (S3, MinIO)

---

# 9️⃣ Характеристики БД

## ACID / BASE

**ACID**
- Atomicity
- Consistency
- Isolation
- Durability

**BASE**
- Basically Available
- Soft state
- Eventually consistent

## OLTP / OLAP / HTAP

- OLTP — транзакции
- OLAP — аналитика
- HTAP — гибрид

## Persistent / In-memory

## Embedded (SQLite)

---

# 🔟 Индексы

## Типы
- B-Tree
- Hash
- Bitmap
- Spatial
- Reversed

## Кластерные / Некластерные

## Селективность

## Функциональные

## Покрывающие

## Разряженные

---

# 1️⃣1️⃣ Транзакции

## WAL
Write-Ahead Logging

## Concurrency
- MVCC
- 2PL

## Уровни изоляции
- READ_UNCOMMITTED
- READ_COMMITTED
- REPEATABLE_READ
- SERIALIZABLE

---

# 1️⃣2️⃣ Паттерны хранения

- Сжатие (LZ4, Zstandard)
- Data cooling
- Batch writes
- Расчёт железа
- Stored procedures
- Materialized views
- Triggers

---

# 1️⃣3️⃣ Очереди сообщений

## Сценарии
- Асинхронная обработка
- Буферизация
- Развязка сервисов

## Гарантии
- At most once
- At least once
- Exactly once

## Retention

---

# 1️⃣4️⃣ Репликация

- Синхронная
- Асинхронная
- Полу-синхронная

## Топологии
- Single leader
- Multi leader
- Leaderless

## Типы
- Логическая
- Физическая

## Failover

## Split Brain

---

# 1️⃣5️⃣ Партиционирование и шардирование

## Партиционирование
- Вертикальное
- Горизонтальное

## Шардирование
- Range-based
- Key-based
- Directory-based

## Rebalancing
- Virtual buckets
- Consistent hashing
- Rendezvous hashing

---

# 1️⃣6️⃣ Распределённое хранение

## CAP
Consistency / Availability / Partition tolerance

## PACELC
Latency vs Consistency

## CDN

## CDC

---

# 1️⃣7️⃣ Паттерны проектирования

## Архитектура
- File-server
- Client-server
- P2P

## Deployment
- Rolling
- Blue-green
- Canary

## Backend
- Монолит
- Микросервисы
- SOA

## Event-Driven
- Event notification
- State transfer
- Event collaboration

## Консенсус
- Raft
- Paxos
- Distributed locking
- SAGA (хореография / оркестрация)
- 2PC / 3PC

## Async Patterns
- Point-to-point
- Pub-sub
- Request-response
- Dead letter queue

---

# 1️⃣8️⃣ Паттерны отказоустойчивости

- Fallback
- Bulkheads
- Backpressure
- Self-healing
- Rate limiting
- Retries (exponential backoff)
- Feature toggles
- Circuit breaker
- Graceful degradation

---

# 🚀 Итог

Этот документ покрывает:
- System Design
- Distributed Systems
- Databases
- Scalability
- Reliability
- Observability
- Fault Tolerance
- Backend Architecture

Подходит для подготовки к:
- FAANG
- MANGA
- Tesla
- SpaceX
- Senior/Staff System Design Interviews
