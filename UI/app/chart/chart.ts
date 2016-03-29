import {Component} from "angular2/core";



@Component({
    selector: 'chart',
    template: `
       chart，
       <p>{{content}}</p>
       
        `,
    styles: [`
       
        .green {
            color: greenyellow;
            }    
    `]


})
export class Chart {
    constructor() {

        this.content = `
证监会新闻发言人邓舸昨天表示，近期，证监会要求所有上市公司制定维护股价稳定的方案，举措包括大股东增持、董监高增持、公司回购、员工持股计划、股权激励等，相关方案应尽快公布，并通过交易所平台向投资者介绍生产经营管理情况、加强投资者关系管理、维护股价稳定、做好投资者沟通工作。他介绍，该措施得到了上市公司大股东的积极响应，包括北京创业板董事长俱乐部、创业板首批28家公司实际控制人、浙江24家公司董事长等多个上市公司联盟以及大连、青岛、湖南等多地上市公司集体发声，宣布通过积极增持、回购、暂不减持等方式稳定公司股价。截至9日，两市已有655家公司公布股份增持、回购计划，积极维护公司股价稳定。
				`;
    }
}


