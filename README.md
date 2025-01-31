https://developers.linear.app/docs
https://rollout.com/integration-guides/linear/sdk/step-by-step-guide-to-building-a-linear-api-integration-in-go

To get `team_id` of linear you have to `CMD + K` and search for `Copy model UUID`


### Todo mvp
- display issues in lipgloss table
- add a "my issues" and "all issues" flag to list issues cmd (default my issues only)
- add a "status" flag to list issues cmd (default is "todo" and "in progress")
    - color those different (todo white and in progress yellow?)
- check config dir (place where we will store env stuff like linear api key)
- create install script



### List of commands supported
- issues
    - [x] list
    - [x] create
    - [ ] update
    - [ ] delete
