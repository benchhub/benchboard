# Benchboard

Frontend for BenchBoard

- [Angular 5](https://angular.io/)
- [Ant Design (ng-zorro)](https://github.com/NG-ZORRO/ng-zorro-antd)
- [ECharts 4](https://github.com/ecomfe/echarts)
  - [ngx-echarts](https://github.com/xieziyu/ngx-echarts)

## Develop

After using angular-cli or paste the latest angular dependencies in [package.json](package.json)

````
npm install --save echarts
npm install --save ngx-echarts
# I never used the storybook ... but just keep there in case we need to prototype some components
npm i -g @storybook/cli
getstorybook
````

## Commands

generated with [Angular CLI](https://github.com/angular/angular-cli) version 1.6.5.

- `ng serve` for a dev server. Navigate to `http://localhost:4200/`. The app will automatically reload if you change any of the source files.
- `ng generate component component-name` to generate a new component. You can also use `ng generate directive|pipe|service|class|guard|interface|enum|module`.
- `ng build` to build the project. The build artifacts will be stored in the `dist/` directory. Use the `-prod` flag for a production build.
- `ng test` to execute the unit tests via [Karma](https://karma-runner.github.io).
- `ng e2e` to execute the end-to-end tests via [Protractor](http://www.protractortest.org/).

To get more help on the Angular CLI use `ng help` or go check out the [Angular CLI README](https://github.com/angular/angular-cli/blob/master/README.md).
