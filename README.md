# Telegram Bot: Article Saver and Reminder
This is a Telegram bot written in Golang that can save articles and remind you to read a random article from your saved articles every 12 hours.

### Features
Save articles: The bot allows you to save articles for later reading.
View your articles: Use command "/view_articles" to view the list of your saved articles.
Reminder: The bot will send you a reminder to read a random article every 12 hours.

### You can use my bot here:
https://t.me/articles_saving_bot

### Installation
Follow these steps to install and run the bot
1. Clone the repository:
```
   git clone https://github.com/sssyrbu/read-later-telegram-bot
   cd read-later-telegram-bot
```
3. Rename the .env.example file to .env and paste your token there.
4. Compose the project using Docker Compose:
``` 
docker-compose up --build
```
5. Run the bot in detached mode:
```
docker-compose up -d
```

### Usage screenshots:
