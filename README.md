# sleeper-go
Go library for the [Sleeper](https://sleeper.com/) fantasy sports free read-only API. According to their [documentation](https://docs.sleeper.com), stay under 1000 API calls per minute, otherwise, you risk being IP-blocked.


## Status of implementations

* [x] User
    * [x] User object - `GET https://api.sleeper.app/v1/user/<username>`
    * [x] User object - `GET https://api.sleeper.app/v1//user/<user_id>`
* [x] Avatars
    * [x] Full size URL - `GET https://sleepercdn.com/avatars/<avatar_id>`
    * [x] Thumbnail URL - `GET https://sleepercdn.com/avatars/thumbs/<avatar_id>`
* [ ] Leagues
    * [x] Get all leagues for user - `GET https://api.sleeper.app/v1/user/<user_id>/leagues/<sport>/<season>`
    * [x] Get a specific league - `GET https://api.sleeper.app/v1/league/<league_id>`
    * [x] Get rosters in a league - `GET https://api.sleeper.app/v1/league/<league_id>/rosters`
    * [x] Get users in a league - `GET https://api.sleeper.app/v1/league/<league_id>/users`
    * [x] Get matchups in a league - `GET https://api.sleeper.app/v1/league/<league_id>/matchups/<week>`
    * [ ] Get the playoff bracket (winners bracket) - `GET https://api.sleeper.app/v1/league/<league_id>/winners_bracket`
    * [ ] Get the playoff bracket (loses bracket) - `GET https://api.sleeper.app/v1/league/<league_id>/loses_bracket`
    * [x] Get transactions - `GET https://api.sleeper.app/v1/league/<league_id>/transactions/<round>`
    * [x] Get traded picks - `GET https://api.sleeper.app/v1/league/<league_id>/traded_picks`
    * [x] Get sport state - `GET https://api.sleeper.app/v1/state/<sport>`
* [x] Drafts
    * [x] Get all drafts for user - `GET https://api.sleeper.app/v1/user/<user_id>/drafts/<sport>/<season>`
    * [x] Get all drafts for a league - `GET https://api.sleeper.app/v1/league/<league_id>/drafts`
    * [x] Get a specific draft - `GET https://api.sleeper.app/v1/draft/<draft_id>`
    * [x] Get all picks in a draft - `GET https://api.sleeper.app/v1/draft/<draft_id>/picks`
    * [x] Get traded picks in a draft - `GET https://api.sleeper.app/v1/draft/<draft_id>/traded_picks`
* [x] Players
    * [x] Get all players - `GET https://api.sleeper.app/v1/players/<sport>`
    * [x] Trending players - `GET https://api.sleeper.app/v1/players/<sport>/trending/<type>?lookback_hours=<hours>&limit=<int>`


## License

MIT
