## <h3>Gonetmaster</h3>

<h6>Command And Control For Ransomware</h6>

---

<p>
    <img src="https://i.pinimg.com/736x/40/a6/e3/40a6e30c124e11047cc7cc913a68567c--astronomy-art-google.jpg" alt="cnc" width="300px"/><br/>
</p>

---

<h4>Introduction</h4>

- Gonetmaster is a command and control for ransomware.
- Includes logger, limiter and much more.
- Easy to customize.

---

<h4>Flow</h4>

Basically, your ransomware will generate a uuid v1 on client side and make a post request to the api endpoint: `/users`. After that, an key to encrypt the files will be generated and store in the database. The key will be received if the insert was sucessful.

The bot master has the possibility to see all the victims at `/users` via a get request with an API token.

The API token can be insert in the database this way:

```sql
INSERT INTO account (key) VALUES ('fdcb0a7c658c0835ab597898462e8f64ce6d87c914217e2a5ce7910f3408699d');
```

The key is a basic uuid v4 turned into base 64 hashed with sha256.

An example of the key can be found in database/database.go inside the function `createTableAccount`.

The get request to this endpoint can be done using curl:

```bash
curl --location '127.0.0.1:3000/users' \
--header 'X-API-KEY: API_KEY'
```

API_KEY is basically the string before hashing.

---

<h4>Warning</h4>

- This project was made for educational purposes only! I take no responsibility for anything you do with this program.
- If you have any suggestions, problems, open a problem (if it is an error, you must be sure to look if you can solve it with [Google](https://giybf.com)!)

<h4>Support me</h4>

- Thanks for looking at this repository, if you like to press the ⭐ button!
- Made by [imzoloft](https://gitlab.com/imzoloft).

<p align="center">
    <b>Informations</b><br>
    <img alt="GitHub Repo stars" src="https://img.shields.io/github/stars/imzoloft/gonetmaster?color=000">
    <img alt="GitHub top language" src="https://img.shields.io/github/languages/top/imzoloft/gonetmaster?color=000">
    <img alt="GitHub last commit" src="https://img.shields.io/github/last-commit/imzoloft/gonetmaster?color=000">
    <img alt="GitHub" src="https://img.shields.io/github/license/imzoloft/gonetmaster?color=000">
    <img alt="GitHub watchers" src="https://img.shields.io/github/watchers/imzoloft/gonetmaster?color=000">
</p>
