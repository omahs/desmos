"use strict";(self.webpackChunkdesmos_docs=self.webpackChunkdesmos_docs||[]).push([[20464],{3905:(e,r,t)=>{t.d(r,{Zo:()=>c,kt:()=>f});var n=t(67294);function a(e,r,t){return r in e?Object.defineProperty(e,r,{value:t,enumerable:!0,configurable:!0,writable:!0}):e[r]=t,e}function o(e,r){var t=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);r&&(n=n.filter((function(r){return Object.getOwnPropertyDescriptor(e,r).enumerable}))),t.push.apply(t,n)}return t}function s(e){for(var r=1;r<arguments.length;r++){var t=null!=arguments[r]?arguments[r]:{};r%2?o(Object(t),!0).forEach((function(r){a(e,r,t[r])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(t)):o(Object(t)).forEach((function(r){Object.defineProperty(e,r,Object.getOwnPropertyDescriptor(t,r))}))}return e}function d(e,r){if(null==e)return{};var t,n,a=function(e,r){if(null==e)return{};var t,n,a={},o=Object.keys(e);for(n=0;n<o.length;n++)t=o[n],r.indexOf(t)>=0||(a[t]=e[t]);return a}(e,r);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(n=0;n<o.length;n++)t=o[n],r.indexOf(t)>=0||Object.prototype.propertyIsEnumerable.call(e,t)&&(a[t]=e[t])}return a}var i=n.createContext({}),l=function(e){var r=n.useContext(i),t=r;return e&&(t="function"==typeof e?e(r):s(s({},r),e)),t},c=function(e){var r=l(e.components);return n.createElement(i.Provider,{value:r},e.children)},p={inlineCode:"code",wrapper:function(e){var r=e.children;return n.createElement(n.Fragment,{},r)}},u=n.forwardRef((function(e,r){var t=e.components,a=e.mdxType,o=e.originalType,i=e.parentName,c=d(e,["components","mdxType","originalType","parentName"]),u=l(t),f=a,m=u["".concat(i,".").concat(f)]||u[f]||p[f]||o;return t?n.createElement(m,s(s({ref:r},c),{},{components:t})):n.createElement(m,s({ref:r},c))}));function f(e,r){var t=arguments,a=r&&r.mdxType;if("string"==typeof e||a){var o=t.length,s=new Array(o);s[0]=u;var d={};for(var i in r)hasOwnProperty.call(r,i)&&(d[i]=r[i]);d.originalType=e,d.mdxType="string"==typeof e?e:a,s[1]=d;for(var l=2;l<o;l++)s[l]=t[l];return n.createElement.apply(null,s)}return n.createElement.apply(null,t)}u.displayName="MDXCreateElement"},48145:(e,r,t)=>{t.r(r),t.d(r,{assets:()=>l,contentTitle:()=>d,default:()=>u,frontMatter:()=>s,metadata:()=>i,toc:()=>c});t(67294);var n=t(3905);function a(){return a=Object.assign||function(e){for(var r=1;r<arguments.length;r++){var t=arguments[r];for(var n in t)Object.prototype.hasOwnProperty.call(t,n)&&(e[n]=t[n])}return e},a.apply(this,arguments)}function o(e,r){if(null==e)return{};var t,n,a=function(e,r){if(null==e)return{};var t,n,a={},o=Object.keys(e);for(n=0;n<o.length;n++)t=o[n],r.indexOf(t)>=0||(a[t]=e[t]);return a}(e,r);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(n=0;n<o.length;n++)t=o[n],r.indexOf(t)>=0||Object.prototype.propertyIsEnumerable.call(e,t)&&(a[t]=e[t])}return a}const s={id:"dtag-transfer-requests-max-order-by",title:"dtag_transfer_requests_max_order_by",hide_table_of_contents:!1},d=void 0,i={unversionedId:"graphql/inputs/dtag-transfer-requests-max-order-by",id:"version-4.2.0/graphql/inputs/dtag-transfer-requests-max-order-by",title:"dtag_transfer_requests_max_order_by",description:'order by max() on columns of table "dtagtransferrequests"',source:"@site/versioned_docs/version-4.2.0/07-graphql/inputs/dtag-transfer-requests-max-order-by.mdx",sourceDirName:"07-graphql/inputs",slug:"/graphql/inputs/dtag-transfer-requests-max-order-by",permalink:"/4.2.0/graphql/inputs/dtag-transfer-requests-max-order-by",draft:!1,editUrl:"https://github.com/desmos-labs/desmos/tree/master/docs/versioned_docs/version-4.2.0/07-graphql/inputs/dtag-transfer-requests-max-order-by.mdx",tags:[],version:"4.2.0",frontMatter:{id:"dtag-transfer-requests-max-order-by",title:"dtag_transfer_requests_max_order_by",hide_table_of_contents:!1},sidebar:"docs",previous:{title:"dtag_transfer_requests_bool_exp",permalink:"/4.2.0/graphql/inputs/dtag-transfer-requests-bool-exp"},next:{title:"dtag_transfer_requests_min_order_by",permalink:"/4.2.0/graphql/inputs/dtag-transfer-requests-min-order-by"}},l={},c=[{value:"Fields",id:"fields",level:3},{value:"<code>receiver_address</code> (<code>order_by</code>)",id:"receiver_address-order_by",level:4},{value:"<code>sender_address</code> (<code>order_by</code>)",id:"sender_address-order_by",level:4}],p={toc:c};function u(e){var{components:r}=e,t=o(e,["components"]);return(0,n.kt)("wrapper",a({},p,t,{components:r,mdxType:"MDXLayout"}),(0,n.kt)("p",null,'order by max() on columns of table "dtag_transfer_requests"'),(0,n.kt)("pre",null,(0,n.kt)("code",a({parentName:"pre"},{className:"language-graphql"}),"input dtag_transfer_requests_max_order_by {\n  receiver_address: order_by\n  sender_address: order_by\n}\n")),(0,n.kt)("h3",a({},{id:"fields"}),"Fields"),(0,n.kt)("h4",a({},{id:"receiver_address-order_by"}),(0,n.kt)("a",a({parentName:"h4"},{href:"#"}),(0,n.kt)("inlineCode",{parentName:"a"},"receiver_address"))," (",(0,n.kt)("a",a({parentName:"h4"},{href:"../enums/order-by"}),(0,n.kt)("inlineCode",{parentName:"a"},"order_by")),")"),(0,n.kt)("h4",a({},{id:"sender_address-order_by"}),(0,n.kt)("a",a({parentName:"h4"},{href:"#"}),(0,n.kt)("inlineCode",{parentName:"a"},"sender_address"))," (",(0,n.kt)("a",a({parentName:"h4"},{href:"../enums/order-by"}),(0,n.kt)("inlineCode",{parentName:"a"},"order_by")),")"))}u.isMDXComponent=!0}}]);