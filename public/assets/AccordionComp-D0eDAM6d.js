import{c as d,_,a as f,o as h,b as m,d as g,e as c,g as a,k as C,t as k,j as w,h as i,u as p,l as x,L as y,E as B}from"./index-oAtpp-VZ.js";/**
 * @license @tabler/icons-vue v3.31.0 - MIT
 *
 * This source code is licensed under the MIT license.
 * See the LICENSE file in the root directory of this source tree.
 */var I=d("outline","chevron-down","IconChevronDown",[["path",{d:"M6 9l6 6l6 -6",key:"svg-0"}]]);/**
 * @license @tabler/icons-vue v3.31.0 - MIT
 *
 * This source code is licensed under the MIT license.
 * See the LICENSE file in the root directory of this source tree.
 */var A=d("outline","chevron-up","IconChevronUp",[["path",{d:"M6 15l6 -6l6 6",key:"svg-0"}]]);const $={class:"accordion"},D={class:"accordion__content"},M={__name:"AccordionComp",props:{title:String,open:Boolean},emits:["onOpen","onClose","onToggle"],setup(l,{expose:u,emit:v}){const s=v;u({toggleAccordion:n});const o=l,e=f(!1);h(()=>{o.open&&n(o.open)}),m(()=>{o.open&&n(o.open)});function n(t=!1){t?(e.value=!0,s("onOpen")):(e.value=!1,s("onClose")),s("onToggle")}return(t,r)=>(c(),g("div",$,[a("header",{onClick:r[0]||(r[0]=O=>n(!e.value))},[a("h4",null,k(l.title),1),C(x,{variant:"ghost",size:"sm",class:"!px-1"},{default:w(()=>[e.value?(c(),i(p(A),{key:1})):(c(),i(p(I),{key:0}))]),_:1})]),a("section",{class:B(`accordion__wrapper ${e.value?"open":""}`)},[a("div",D,[y(t.$slots,"default",{},void 0,!0)])],2)]))}},U=_(M,[["__scopeId","data-v-0c582c8a"]]);export{U as A};
//# sourceMappingURL=AccordionComp-D0eDAM6d.js.map
