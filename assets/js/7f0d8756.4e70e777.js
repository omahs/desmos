"use strict";(self.webpackChunkdesmos_docs=self.webpackChunkdesmos_docs||[]).push([[51118],{3905:(e,r,t)=>{t.d(r,{Zo:()=>i,kt:()=>b});var o=t(67294);function n(e,r,t){return r in e?Object.defineProperty(e,r,{value:t,enumerable:!0,configurable:!0,writable:!0}):e[r]=t,e}function s(e,r){var t=Object.keys(e);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);r&&(o=o.filter((function(r){return Object.getOwnPropertyDescriptor(e,r).enumerable}))),t.push.apply(t,o)}return t}function p(e){for(var r=1;r<arguments.length;r++){var t=null!=arguments[r]?arguments[r]:{};r%2?s(Object(t),!0).forEach((function(r){n(e,r,t[r])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(t)):s(Object(t)).forEach((function(r){Object.defineProperty(e,r,Object.getOwnPropertyDescriptor(t,r))}))}return e}function a(e,r){if(null==e)return{};var t,o,n=function(e,r){if(null==e)return{};var t,o,n={},s=Object.keys(e);for(o=0;o<s.length;o++)t=s[o],r.indexOf(t)>=0||(n[t]=e[t]);return n}(e,r);if(Object.getOwnPropertySymbols){var s=Object.getOwnPropertySymbols(e);for(o=0;o<s.length;o++)t=s[o],r.indexOf(t)>=0||Object.prototype.propertyIsEnumerable.call(e,t)&&(n[t]=e[t])}return n}var d=o.createContext({}),u=function(e){var r=o.useContext(d),t=r;return e&&(t="function"==typeof e?e(r):p(p({},r),e)),t},i=function(e){var r=u(e.components);return o.createElement(d.Provider,{value:r},e.children)},c={inlineCode:"code",wrapper:function(e){var r=e.children;return o.createElement(o.Fragment,{},r)}},l=o.forwardRef((function(e,r){var t=e.components,n=e.mdxType,s=e.originalType,d=e.parentName,i=a(e,["components","mdxType","originalType","parentName"]),l=u(t),b=n,y=l["".concat(d,".").concat(b)]||l[b]||c[b]||s;return t?o.createElement(y,p(p({ref:r},i),{},{components:t})):o.createElement(y,p({ref:r},i))}));function b(e,r){var t=arguments,n=r&&r.mdxType;if("string"==typeof e||n){var s=t.length,p=new Array(s);p[0]=l;var a={};for(var d in r)hasOwnProperty.call(r,d)&&(a[d]=r[d]);a.originalType=e,a.mdxType="string"==typeof e?e:n,p[1]=a;for(var u=2;u<s;u++)p[u]=t[u];return o.createElement.apply(null,p)}return o.createElement.apply(null,t)}l.displayName="MDXCreateElement"},93562:(e,r,t)=>{t.r(r),t.d(r,{assets:()=>u,contentTitle:()=>a,default:()=>l,frontMatter:()=>p,metadata:()=>d,toc:()=>i});t(67294);var o=t(3905);function n(){return n=Object.assign||function(e){for(var r=1;r<arguments.length;r++){var t=arguments[r];for(var o in t)Object.prototype.hasOwnProperty.call(t,o)&&(e[o]=t[o])}return e},n.apply(this,arguments)}function s(e,r){if(null==e)return{};var t,o,n=function(e,r){if(null==e)return{};var t,o,n={},s=Object.keys(e);for(o=0;o<s.length;o++)t=s[o],r.indexOf(t)>=0||(n[t]=e[t]);return n}(e,r);if(Object.getOwnPropertySymbols){var s=Object.getOwnPropertySymbols(e);for(o=0;o<s.length;o++)t=s[o],r.indexOf(t)>=0||Object.prototype.propertyIsEnumerable.call(e,t)&&(n[t]=e[t])}return n}const p={id:"subspace-user-group-stddev-pop-order-by",title:"subspace_user_group_stddev_pop_order_by",hide_table_of_contents:!1},a=void 0,d={unversionedId:"graphql/inputs/subspace-user-group-stddev-pop-order-by",id:"version-4.2.0/graphql/inputs/subspace-user-group-stddev-pop-order-by",title:"subspace_user_group_stddev_pop_order_by",description:'order by stddevpop() on columns of table "subspaceuser_group"',source:"@site/versioned_docs/version-4.2.0/07-graphql/inputs/subspace-user-group-stddev-pop-order-by.mdx",sourceDirName:"07-graphql/inputs",slug:"/graphql/inputs/subspace-user-group-stddev-pop-order-by",permalink:"/4.2.0/graphql/inputs/subspace-user-group-stddev-pop-order-by",draft:!1,editUrl:"https://github.com/desmos-labs/desmos/tree/master/docs/versioned_docs/version-4.2.0/07-graphql/inputs/subspace-user-group-stddev-pop-order-by.mdx",tags:[],version:"4.2.0",frontMatter:{id:"subspace-user-group-stddev-pop-order-by",title:"subspace_user_group_stddev_pop_order_by",hide_table_of_contents:!1},sidebar:"docs",previous:{title:"subspace_user_group_stddev_order_by",permalink:"/4.2.0/graphql/inputs/subspace-user-group-stddev-order-by"},next:{title:"subspace_user_group_stddev_samp_order_by",permalink:"/4.2.0/graphql/inputs/subspace-user-group-stddev-samp-order-by"}},u={},i=[{value:"Fields",id:"fields",level:3},{value:"<code>id</code> (<code>order_by</code>)",id:"id-order_by",level:4},{value:"<code>subspace_id</code> (<code>order_by</code>)",id:"subspace_id-order_by",level:4}],c={toc:i};function l(e){var{components:r}=e,t=s(e,["components"]);return(0,o.kt)("wrapper",n({},c,t,{components:r,mdxType:"MDXLayout"}),(0,o.kt)("p",null,'order by stddev_pop() on columns of table "subspace_user_group"'),(0,o.kt)("pre",null,(0,o.kt)("code",n({parentName:"pre"},{className:"language-graphql"}),"input subspace_user_group_stddev_pop_order_by {\n  id: order_by\n  subspace_id: order_by\n}\n")),(0,o.kt)("h3",n({},{id:"fields"}),"Fields"),(0,o.kt)("h4",n({},{id:"id-order_by"}),(0,o.kt)("a",n({parentName:"h4"},{href:"#"}),(0,o.kt)("inlineCode",{parentName:"a"},"id"))," (",(0,o.kt)("a",n({parentName:"h4"},{href:"../enums/order-by"}),(0,o.kt)("inlineCode",{parentName:"a"},"order_by")),")"),(0,o.kt)("h4",n({},{id:"subspace_id-order_by"}),(0,o.kt)("a",n({parentName:"h4"},{href:"#"}),(0,o.kt)("inlineCode",{parentName:"a"},"subspace_id"))," (",(0,o.kt)("a",n({parentName:"h4"},{href:"../enums/order-by"}),(0,o.kt)("inlineCode",{parentName:"a"},"order_by")),")"))}l.isMDXComponent=!0}}]);