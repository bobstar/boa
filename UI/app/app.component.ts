import {Component} from 'angular2/core';
import {Topbar} from './topbar/topbar';
import {Chart} from './chart/chart';

@Component({
    selector: 'my-app',
    template: `
    <topbar></topbar>
    <hr>
    <chart></chart>
    `,
    directives:[Topbar,Chart]
})
export  class AppComponent { }
