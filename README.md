# ReportPortal Demo project

A demo project to test [reportportal.io](https://github.com/reportportal/reportportal). 

## Start

To start the demo run: `docker compose up -d`

Login with user: `default` and password: `1q2w3e`. 

Following steps create a test-report:
- create a dashboard
- create a filter
    - cnt: report
- add a widget to the dashboard
    - select the widget type
    - select the filter

## Jest

To run and upload tests run `npm test`. 

Jest will run all tests and `reportportal/agent-js-jest` will upload 

## Go

To run and upload tests run `sh test.sh`.
