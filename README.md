# Tic-Tac-Toe Online

Многопользовательская онлайн игра в крестики-нолики с поддержкой быстрого поиска противников, режима оффлайн и real-time коммуникацией через WebSocket.

## 🚀 Возможности

- **Быстрая игра**: поиск случайного противника онлайн
- **Оффлайн режим**: локальная игра для двух игроков
- **Real-time обновления**: мгновенная синхронизация ходов через WebSocket
- **Автоматические никнеймы**: уникальные случайные имена для каждого игрока
- **Статистика**: отображение активных игроков и игр
- **Адаптивный дизайн**: работает на всех устройствах

## 🛠️ Технологии

### Frontend
- **HTML5/CSS3**: современный адаптивный интерфейс
- **JavaScript**: чистый JS без фреймворков
- **WebSocket API**: real-time коммуникация

### Backend
- **Node.js** или **Go**: серверная логика
- **WebSocket**: двусторонняя коммуникация
- **Long Polling**: обновление статистики
- **HTTP REST API**: управление игровыми сессиями

## 📦 Установка

### Требования

- Node.js 16+ (или Go 1.21+)
- Современный браузер с поддержкой WebSocket

### Запуск проекта

```bash
# Клонировать репозиторий
git clone https://github.com/ebairamo/tic-tac-toe.git
cd tic-tac-toe

# Установить зависимости (для Node.js версии)
npm install

# Запустить сервер
npm start

# Или для Go версии
go run main.go
```

Откройте браузер: `http://localhost:3000`

## 🎮 Как играть

### Быстрая игра

1. На главной странице нажмите **"Quick Game"**
2. Система найдет вам противника из пула ожидающих игроков
3. При нажатии **"Cancel"** поиск отменяется
4. После нахождения противника игра начинается автоматически
5. Символы (X или O) назначаются случайным образом
6. Игроки ходят по очереди
7. После завершения доступны кнопки:
   - **"Play Again"** - реванш с тем же противником
   - **"Back to Main Menu"** - вернуться на главную

### Оффлайн режим

1. Нажмите **"Offline Game"** на главной странице
2. Играйте локально с другом на одном устройстве
3. Игра автоматически начинается заново после завершения
4. Отслеживается простая статистика (победы/поражения/ничьи)
5. Кнопка **"Back to Main"** доступна всегда

## 🏗️ Архитектура

### Коммуникация

```
┌─────────────┐                    ┌─────────────┐
│   Player 1  │◄──────WebSocket────►│   Server    │
│  (Browser)  │                     │             │
└─────────────┘                     └──────┬──────┘
                                           │
                                    WebSocket
                                           │
┌─────────────┐                     ┌──────▼──────┐
│   Player 2  │◄──────WebSocket─────┤  Match Pool │
│  (Browser)  │                     └─────────────┘
└─────────────┘
```

### Поток игры

**Фаза 1: Поиск игры**
1. Игрок нажимает "Quick Game"
2. Добавляется в пул ожидающих
3. Сервер подбирает пару из пула
4. Оба игрока получают уведомление

**Фаза 2: Геймплей**
1. Случайное назначение X и O
2. Игроки ходят по очереди
3. Каждый ход синхронизируется через WebSocket
4. Проверка победных комбинаций

**Фаза 3: Завершение**
1. Определение победителя/ничьей
2. Отображение результата
3. Опции: реванш или выход в меню

## 📊 Особенности реализации

### WebSocket протокол

**Клиент → Сервер:**
```json
{
  "type": "move",
  "position": 4,
  "gameId": "abc123"
}
```

**Сервер → Клиент:**
```json
{
  "type": "game_update",
  "board": [null, "X", null, "O", "X", null, null, null, null],
  "currentPlayer": "O",
  "status": "in_progress"
}
```

### Long Polling для статистики

Обновление каждые 60 секунд:
- Количество игроков онлайн
- Количество активных игр

### Обработка отключений

- **Игрок отключился**: противник получает уведомление, игра завершается
- **Переподключение**: восстановление сессии (опционально)
- **Таймауты**: автоматическое завершение при долгом бездействии

## 🎨 Интерфейс

### Главная страница
- Кнопки режимов игры
- Live статистика (игроки/игры)
- Случайный никнейм

### Игровая комната
- 3×3 игровое поле
- Индикатор текущего игрока
- Информация о противнике
- Кнопки управления

### Экран результата
- Сообщение о победе/поражении/ничьей
- Счет игры
- Опции продолжения

## 🔧 Разработка

### Структура проекта

```
tic-tac-toe/
├── public/
│   ├── index.html          # Главная страница
│   ├── game.html           # Игровая комната
│   ├── style.css           # Стили
│   └── script.js           # Клиентская логика
├── server/
│   ├── main.js             # Сервер (Node.js)
│   ├── game.js             # Игровая логика
│   └── matchmaking.js      # Поиск противников
└── README.md
```

### Методы HTTP коммуникации

- **WebSocket**: real-time геймплей
- **Long Polling**: статистика онлайн игроков
- **HTTP GET/POST**: создание/завершение игр
- **SSE** (опционально): режим наблюдателя

## ✅ Валидация

- **W3C Validator**: HTML без ошибок
- **ESLint**: следование правилам:
  ```json
  {
    "semi": "error",
    "no-unused-vars": "error",
    "no-var": "error",
    "no-undef": "error"
  }
  ```

## 🎓 Цели обучения

Этот проект демонстрирует:
- Клиент-серверную коммуникацию
- WebSocket для real-time приложений
- Управление состоянием игры
- Matchmaking алгоритмы
- HTTP методы и паттерны
- Обработку подключений/отключений
- Адаптивный веб-дизайн

## 🙏 Автор задания

**Dias Kappassov** - Software Developer at Doodocs.kz
- Email: diaskappassov@gmail.com
- [GitHub](https://github.com/Dias1c)
- [LinkedIn](https://www.linkedin.com/in/dias-kappassov/)

---

*Проект выполнен в рамках обучения в ALEM School*

# Tic-Tac-Toe Online

Multiplayer online Tic-Tac-Toe game with quick opponent matching, offline mode, and real-time WebSocket communication.

## 🚀 Features

- **Quick Game**: find random opponent online
- **Offline Mode**: local two-player game
- **Real-time Updates**: instant move synchronization via WebSocket
- **Automatic Nicknames**: unique random names for each player
- **Statistics**: display of active players and games
- **Responsive Design**: works on all devices

## 🛠️ Tech Stack

### Frontend
- **HTML5/CSS3**: modern responsive interface
- **JavaScript**: vanilla JS without frameworks
- **WebSocket API**: real-time communication

### Backend
- **Node.js** or **Go**: server logic
- **WebSocket**: bidirectional communication
- **Long Polling**: statistics updates
- **HTTP REST API**: game session management

## 📦 Installation

### Requirements

- Node.js 16+ (or Go 1.21+)
- Modern browser with WebSocket support

### Running the Project

```bash
# Clone repository
git clone https://github.com/ebairamo/tic-tac-toe.git
cd tic-tac-toe

# Install dependencies (for Node.js version)
npm install

# Start server
npm start

# Or for Go version
go run main.go
```

Open browser: `http://localhost:3000`

## 🎮 How to Play

### Quick Game

1. On the main page, click **"Quick Game"**
2. System finds you an opponent from the waiting pool
3. Click **"Cancel"** to abort search
4. After finding opponent, game starts automatically
5. Symbols (X or O) assigned randomly
6. Players take turns
7. After completion, buttons available:
   - **"Play Again"** - rematch with same opponent
   - **"Back to Main Menu"** - return to main page

### Offline Mode

1. Click **"Offline Game"** on main page
2. Play locally with a friend on same device
3. Game automatically restarts after completion
4. Simple statistics tracked (wins/losses/draws)
5. **"Back to Main"** button always available

## 🏗️ Architecture

### Communication

```
┌─────────────┐                    ┌─────────────┐
│   Player 1  │◄──────WebSocket────►│   Server    │
│  (Browser)  │                     │             │
└─────────────┘                     └──────┬──────┘
                                           │
                                    WebSocket
                                           │
┌─────────────┐                     ┌──────▼──────┐
│   Player 2  │◄──────WebSocket─────┤  Match Pool │
│  (Browser)  │                     └─────────────┘
└─────────────┘
```

### Game Flow

**Phase 1: Game Search**
1. Player clicks "Quick Game"
2. Added to waiting pool
3. Server matches pair from pool
4. Both players receive notification

**Phase 2: Gameplay**
1. Random X and O assignment
2. Players take turns
3. Each move synced via WebSocket
4. Win condition checking

**Phase 3: Completion**
1. Winner/draw determination
2. Result display
3. Options: rematch or exit to menu

## 📊 Implementation Features

### WebSocket Protocol

**Client → Server:**
```json
{
  "type": "move",
  "position": 4,
  "gameId": "abc123"
}
```

**Server → Client:**
```json
{
  "type": "game_update",
  "board": [null, "X", null, "O", "X", null, null, null, null],
  "currentPlayer": "O",
  "status": "in_progress"
}
```

### Long Polling for Statistics

Updates every 60 seconds:
- Number of online players
- Number of active games

### Disconnection Handling

- **Player disconnected**: opponent receives notification, game ends
- **Reconnection**: session restoration (optional)
- **Timeouts**: automatic completion on long inactivity

## 🎨 Interface

### Main Page
- Game mode buttons
- Live statistics (players/games)
- Random nickname

### Game Room
- 3×3 game board
- Current player indicator
- Opponent information
- Control buttons

### Result Screen
- Win/loss/draw message
- Game score
- Continue options

## 🔧 Development

### Project Structure

```
tic-tac-toe/
├── public/
│   ├── index.html          # Main page
│   ├── game.html           # Game room
│   ├── style.css           # Styles
│   └── script.js           # Client logic
├── server/
│   ├── main.js             # Server (Node.js)
│   ├── game.js             # Game logic
│   └── matchmaking.js      # Opponent matching
└── README.md
```

### HTTP Communication Methods

- **WebSocket**: real-time gameplay
- **Long Polling**: online player statistics
- **HTTP GET/POST**: game creation/completion
- **SSE** (optional): observer mode

## ✅ Validation

- **W3C Validator**: error-free HTML
- **ESLint**: following rules:
  ```json
  {
    "semi": "error",
    "no-unused-vars": "error",
    "no-var": "error",
    "no-undef": "error"
  }
  ```

## 🎓 Learning Objectives

This project demonstrates:
- Client-server communication
- WebSocket for real-time applications
- Game state management
- Matchmaking algorithms
- HTTP methods and patterns
- Connection/disconnection handling
- Responsive web design

## 🙏 Project Author

**Dias Kappassov** - Software Developer at Doodocs.kz
- Email: diaskappassov@gmail.com
- [GitHub](https://github.com/Dias1c)
- [LinkedIn](https://www.linkedin.com/in/dias-kappassov/)

---

*Project completed as part of ALEM School curriculum*
