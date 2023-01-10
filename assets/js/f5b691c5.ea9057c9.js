"use strict";(self.webpackChunkdesmos_docs=self.webpackChunkdesmos_docs||[]).push([[63837],{3905:(e,r,n)=>{n.d(r,{Zo:()=>u,kt:()=>f});var t=n(67294);function l(e,r,n){return r in e?Object.defineProperty(e,r,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[r]=n,e}function s(e,r){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var t=Object.getOwnPropertySymbols(e);r&&(t=t.filter((function(r){return Object.getOwnPropertyDescriptor(e,r).enumerable}))),n.push.apply(n,t)}return n}function a(e){for(var r=1;r<arguments.length;r++){var n=null!=arguments[r]?arguments[r]:{};r%2?s(Object(n),!0).forEach((function(r){l(e,r,n[r])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):s(Object(n)).forEach((function(r){Object.defineProperty(e,r,Object.getOwnPropertyDescriptor(n,r))}))}return e}function o(e,r){if(null==e)return{};var n,t,l=function(e,r){if(null==e)return{};var n,t,l={},s=Object.keys(e);for(t=0;t<s.length;t++)n=s[t],r.indexOf(n)>=0||(l[n]=e[n]);return l}(e,r);if(Object.getOwnPropertySymbols){var s=Object.getOwnPropertySymbols(e);for(t=0;t<s.length;t++)n=s[t],r.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(l[n]=e[n])}return l}var c=t.createContext({}),i=function(e){var r=t.useContext(c),n=r;return e&&(n="function"==typeof e?e(r):a(a({},r),e)),n},u=function(e){var r=i(e.components);return t.createElement(c.Provider,{value:r},e.children)},d={inlineCode:"code",wrapper:function(e){var r=e.children;return t.createElement(t.Fragment,{},r)}},p=t.forwardRef((function(e,r){var n=e.components,l=e.mdxType,s=e.originalType,c=e.parentName,u=o(e,["components","mdxType","originalType","parentName"]),p=i(n),f=l,m=p["".concat(c,".").concat(f)]||p[f]||d[f]||s;return n?t.createElement(m,a(a({ref:r},u),{},{components:n})):t.createElement(m,a({ref:r},u))}));function f(e,r){var n=arguments,l=r&&r.mdxType;if("string"==typeof e||l){var s=n.length,a=new Array(s);a[0]=p;var o={};for(var c in r)hasOwnProperty.call(r,c)&&(o[c]=r[c]);o.originalType=e,o.mdxType="string"==typeof e?e:l,a[1]=o;for(var i=2;i<s;i++)a[i]=n[i];return t.createElement.apply(null,a)}return t.createElement.apply(null,n)}p.displayName="MDXCreateElement"},6242:(e,r,n)=>{n.r(r),n.d(r,{assets:()=>i,contentTitle:()=>o,default:()=>p,frontMatter:()=>a,metadata:()=>c,toc:()=>u});n(67294);var t=n(3905);function l(){return l=Object.assign||function(e){for(var r=1;r<arguments.length;r++){var n=arguments[r];for(var t in n)Object.prototype.hasOwnProperty.call(n,t)&&(e[t]=n[t])}return e},l.apply(this,arguments)}function s(e,r){if(null==e)return{};var n,t,l=function(e,r){if(null==e)return{};var n,t,l={},s=Object.keys(e);for(t=0;t<s.length;t++)n=s[t],r.indexOf(n)>=0||(l[n]=e[n]);return l}(e,r);if(Object.getOwnPropertySymbols){var s=Object.getOwnPropertySymbols(e);for(t=0;t<s.length;t++)n=s[t],r.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(l[n]=e[n])}return l}const a={id:"order-by",title:"order_by",hide_table_of_contents:!1},o=void 0,c={unversionedId:"graphql/enums/order-by",id:"version-4.2.0/graphql/enums/order-by",title:"order_by",description:"column ordering options",source:"@site/versioned_docs/version-4.2.0/07-graphql/enums/order-by.mdx",sourceDirName:"07-graphql/enums",slug:"/graphql/enums/order-by",permalink:"/4.2.0/graphql/enums/order-by",draft:!1,editUrl:"https://github.com/desmos-labs/desmos/tree/master/docs/versioned_docs/version-4.2.0/07-graphql/enums/order-by.mdx",tags:[],version:"4.2.0",frontMatter:{id:"order-by",title:"order_by",hide_table_of_contents:!1},sidebar:"docs",previous:{title:"dtag_transfer_requests_select_column",permalink:"/4.2.0/graphql/enums/dtag-transfer-requests-select-column"},next:{title:"poll_answer_select_column",permalink:"/4.2.0/graphql/enums/poll-answer-select-column"}},i={},u=[{value:"Values",id:"values",level:3},{value:"<code>asc</code>",id:"asc",level:4},{value:"<code>asc_nulls_first</code>",id:"asc_nulls_first",level:4},{value:"<code>asc_nulls_last</code>",id:"asc_nulls_last",level:4},{value:"<code>desc</code>",id:"desc",level:4},{value:"<code>desc_nulls_first</code>",id:"desc_nulls_first",level:4},{value:"<code>desc_nulls_last</code>",id:"desc_nulls_last",level:4}],d={toc:u};function p(e){var{components:r}=e,n=s(e,["components"]);return(0,t.kt)("wrapper",l({},d,n,{components:r,mdxType:"MDXLayout"}),(0,t.kt)("p",null,"column ordering options"),(0,t.kt)("pre",null,(0,t.kt)("code",l({parentName:"pre"},{className:"language-graphql"}),"enum order_by {\n  asc\n  asc_nulls_first\n  asc_nulls_last\n  desc\n  desc_nulls_first\n  desc_nulls_last\n}\n")),(0,t.kt)("h3",l({},{id:"values"}),"Values"),(0,t.kt)("h4",l({},{id:"asc"}),(0,t.kt)("a",l({parentName:"h4"},{href:"#"}),(0,t.kt)("inlineCode",{parentName:"a"},"asc"))),(0,t.kt)("p",null,"in ascending order, nulls last"),(0,t.kt)("h4",l({},{id:"asc_nulls_first"}),(0,t.kt)("a",l({parentName:"h4"},{href:"#"}),(0,t.kt)("inlineCode",{parentName:"a"},"asc_nulls_first"))),(0,t.kt)("p",null,"in ascending order, nulls first"),(0,t.kt)("h4",l({},{id:"asc_nulls_last"}),(0,t.kt)("a",l({parentName:"h4"},{href:"#"}),(0,t.kt)("inlineCode",{parentName:"a"},"asc_nulls_last"))),(0,t.kt)("p",null,"in ascending order, nulls last"),(0,t.kt)("h4",l({},{id:"desc"}),(0,t.kt)("a",l({parentName:"h4"},{href:"#"}),(0,t.kt)("inlineCode",{parentName:"a"},"desc"))),(0,t.kt)("p",null,"in descending order, nulls first"),(0,t.kt)("h4",l({},{id:"desc_nulls_first"}),(0,t.kt)("a",l({parentName:"h4"},{href:"#"}),(0,t.kt)("inlineCode",{parentName:"a"},"desc_nulls_first"))),(0,t.kt)("p",null,"in descending order, nulls first"),(0,t.kt)("h4",l({},{id:"desc_nulls_last"}),(0,t.kt)("a",l({parentName:"h4"},{href:"#"}),(0,t.kt)("inlineCode",{parentName:"a"},"desc_nulls_last"))),(0,t.kt)("p",null,"in descending order, nulls last"))}p.isMDXComponent=!0}}]);