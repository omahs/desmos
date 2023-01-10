"use strict";(self.webpackChunkdesmos_docs=self.webpackChunkdesmos_docs||[]).push([[69954],{3905:(e,r,t)=>{t.d(r,{Zo:()=>d,kt:()=>m});var n=t(67294);function o(e,r,t){return r in e?Object.defineProperty(e,r,{value:t,enumerable:!0,configurable:!0,writable:!0}):e[r]=t,e}function a(e,r){var t=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);r&&(n=n.filter((function(r){return Object.getOwnPropertyDescriptor(e,r).enumerable}))),t.push.apply(t,n)}return t}function i(e){for(var r=1;r<arguments.length;r++){var t=null!=arguments[r]?arguments[r]:{};r%2?a(Object(t),!0).forEach((function(r){o(e,r,t[r])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(t)):a(Object(t)).forEach((function(r){Object.defineProperty(e,r,Object.getOwnPropertyDescriptor(t,r))}))}return e}function p(e,r){if(null==e)return{};var t,n,o=function(e,r){if(null==e)return{};var t,n,o={},a=Object.keys(e);for(n=0;n<a.length;n++)t=a[n],r.indexOf(t)>=0||(o[t]=e[t]);return o}(e,r);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);for(n=0;n<a.length;n++)t=a[n],r.indexOf(t)>=0||Object.prototype.propertyIsEnumerable.call(e,t)&&(o[t]=e[t])}return o}var s=n.createContext({}),c=function(e){var r=n.useContext(s),t=r;return e&&(t="function"==typeof e?e(r):i(i({},r),e)),t},d=function(e){var r=c(e.components);return n.createElement(s.Provider,{value:r},e.children)},l={inlineCode:"code",wrapper:function(e){var r=e.children;return n.createElement(n.Fragment,{},r)}},u=n.forwardRef((function(e,r){var t=e.components,o=e.mdxType,a=e.originalType,s=e.parentName,d=p(e,["components","mdxType","originalType","parentName"]),u=c(t),m=o,y=u["".concat(s,".").concat(m)]||u[m]||l[m]||a;return t?n.createElement(y,i(i({ref:r},d),{},{components:t})):n.createElement(y,i({ref:r},d))}));function m(e,r){var t=arguments,o=r&&r.mdxType;if("string"==typeof e||o){var a=t.length,i=new Array(a);i[0]=u;var p={};for(var s in r)hasOwnProperty.call(r,s)&&(p[s]=r[s]);p.originalType=e,p.mdxType="string"==typeof e?e:o,i[1]=p;for(var c=2;c<a;c++)i[c]=t[c];return n.createElement.apply(null,i)}return n.createElement.apply(null,t)}u.displayName="MDXCreateElement"},43330:(e,r,t)=>{t.r(r),t.d(r,{assets:()=>c,contentTitle:()=>p,default:()=>u,frontMatter:()=>i,metadata:()=>s,toc:()=>d});t(67294);var n=t(3905);function o(){return o=Object.assign||function(e){for(var r=1;r<arguments.length;r++){var t=arguments[r];for(var n in t)Object.prototype.hasOwnProperty.call(t,n)&&(e[n]=t[n])}return e},o.apply(this,arguments)}function a(e,r){if(null==e)return{};var t,n,o=function(e,r){if(null==e)return{};var t,n,o={},a=Object.keys(e);for(n=0;n<a.length;n++)t=a[n],r.indexOf(t)>=0||(o[t]=e[t]);return o}(e,r);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);for(n=0;n<a.length;n++)t=a[n],r.indexOf(t)>=0||Object.prototype.propertyIsEnumerable.call(e,t)&&(o[t]=e[t])}return o}const i={id:"post-mention-variance-order-by",title:"post_mention_variance_order_by",hide_table_of_contents:!1},p=void 0,s={unversionedId:"graphql/inputs/post-mention-variance-order-by",id:"version-4.2.0/graphql/inputs/post-mention-variance-order-by",title:"post_mention_variance_order_by",description:'order by variance() on columns of table "post_mention"',source:"@site/versioned_docs/version-4.2.0/07-graphql/inputs/post-mention-variance-order-by.mdx",sourceDirName:"07-graphql/inputs",slug:"/graphql/inputs/post-mention-variance-order-by",permalink:"/4.2.0/graphql/inputs/post-mention-variance-order-by",draft:!1,editUrl:"https://github.com/desmos-labs/desmos/tree/master/docs/versioned_docs/version-4.2.0/07-graphql/inputs/post-mention-variance-order-by.mdx",tags:[],version:"4.2.0",frontMatter:{id:"post-mention-variance-order-by",title:"post_mention_variance_order_by",hide_table_of_contents:!1},sidebar:"docs",previous:{title:"post_mention_var_samp_order_by",permalink:"/4.2.0/graphql/inputs/post-mention-var-samp-order-by"},next:{title:"post_min_order_by",permalink:"/4.2.0/graphql/inputs/post-min-order-by"}},c={},d=[{value:"Fields",id:"fields",level:3},{value:"<code>end_index</code> (<code>order_by</code>)",id:"end_index-order_by",level:4},{value:"<code>start_index</code> (<code>order_by</code>)",id:"start_index-order_by",level:4}],l={toc:d};function u(e){var{components:r}=e,t=a(e,["components"]);return(0,n.kt)("wrapper",o({},l,t,{components:r,mdxType:"MDXLayout"}),(0,n.kt)("p",null,'order by variance() on columns of table "post_mention"'),(0,n.kt)("pre",null,(0,n.kt)("code",o({parentName:"pre"},{className:"language-graphql"}),"input post_mention_variance_order_by {\n  end_index: order_by\n  start_index: order_by\n}\n")),(0,n.kt)("h3",o({},{id:"fields"}),"Fields"),(0,n.kt)("h4",o({},{id:"end_index-order_by"}),(0,n.kt)("a",o({parentName:"h4"},{href:"#"}),(0,n.kt)("inlineCode",{parentName:"a"},"end_index"))," (",(0,n.kt)("a",o({parentName:"h4"},{href:"../enums/order-by"}),(0,n.kt)("inlineCode",{parentName:"a"},"order_by")),")"),(0,n.kt)("h4",o({},{id:"start_index-order_by"}),(0,n.kt)("a",o({parentName:"h4"},{href:"#"}),(0,n.kt)("inlineCode",{parentName:"a"},"start_index"))," (",(0,n.kt)("a",o({parentName:"h4"},{href:"../enums/order-by"}),(0,n.kt)("inlineCode",{parentName:"a"},"order_by")),")"))}u.isMDXComponent=!0}}]);