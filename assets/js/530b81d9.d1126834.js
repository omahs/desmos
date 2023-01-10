"use strict";(self.webpackChunkdesmos_docs=self.webpackChunkdesmos_docs||[]).push([[4782],{3905:(e,t,r)=>{r.d(t,{Zo:()=>c,kt:()=>u});var n=r(67294);function o(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function s(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,n)}return r}function a(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?s(Object(r),!0).forEach((function(t){o(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):s(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}function i(e,t){if(null==e)return{};var r,n,o=function(e,t){if(null==e)return{};var r,n,o={},s=Object.keys(e);for(n=0;n<s.length;n++)r=s[n],t.indexOf(r)>=0||(o[r]=e[r]);return o}(e,t);if(Object.getOwnPropertySymbols){var s=Object.getOwnPropertySymbols(e);for(n=0;n<s.length;n++)r=s[n],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(o[r]=e[r])}return o}var l=n.createContext({}),p=function(e){var t=n.useContext(l),r=t;return e&&(r="function"==typeof e?e(t):a(a({},t),e)),r},c=function(e){var t=p(e.components);return n.createElement(l.Provider,{value:t},e.children)},m={inlineCode:"code",wrapper:function(e){var t=e.children;return n.createElement(n.Fragment,{},t)}},d=n.forwardRef((function(e,t){var r=e.components,o=e.mdxType,s=e.originalType,l=e.parentName,c=i(e,["components","mdxType","originalType","parentName"]),d=p(r),u=o,f=d["".concat(l,".").concat(u)]||d[u]||m[u]||s;return r?n.createElement(f,a(a({ref:t},c),{},{components:r})):n.createElement(f,a({ref:t},c))}));function u(e,t){var r=arguments,o=t&&t.mdxType;if("string"==typeof e||o){var s=r.length,a=new Array(s);a[0]=d;var i={};for(var l in t)hasOwnProperty.call(t,l)&&(i[l]=t[l]);i.originalType=e,i.mdxType="string"==typeof e?e:o,a[1]=i;for(var p=2;p<s;p++)a[p]=r[p];return n.createElement.apply(null,a)}return n.createElement.apply(null,r)}d.displayName="MDXCreateElement"},69394:(e,t,r)=>{r.r(t),r.d(t,{assets:()=>p,contentTitle:()=>i,default:()=>d,frontMatter:()=>a,metadata:()=>l,toc:()=>c});r(67294);var n=r(3905);function o(){return o=Object.assign||function(e){for(var t=1;t<arguments.length;t++){var r=arguments[t];for(var n in r)Object.prototype.hasOwnProperty.call(r,n)&&(e[n]=r[n])}return e},o.apply(this,arguments)}function s(e,t){if(null==e)return{};var r,n,o=function(e,t){if(null==e)return{};var r,n,o={},s=Object.keys(e);for(n=0;n<s.length;n++)r=s[n],t.indexOf(r)>=0||(o[r]=e[r]);return o}(e,t);if(Object.getOwnPropertySymbols){var s=Object.getOwnPropertySymbols(e);for(n=0;n<s.length;n++)r=s[n],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(o[r]=e[r])}return o}const a={id:"permissions",title:"Permissions",sidebar_label:"Permissions",slug:"permissions"},i="Permissions",l={unversionedId:"developers/modules/reactions/permissions",id:"developers/modules/reactions/permissions",title:"Permissions",description:"Based on which x/reactions related action your users want to perform, they need to be granted with one or more",source:"@site/docs/02-developers/02-modules/reactions/06-permissions.md",sourceDirName:"02-developers/02-modules/reactions",slug:"/developers/modules/reactions/permissions",permalink:"/developers/modules/reactions/permissions",draft:!1,editUrl:"https://github.com/desmos-labs/desmos/tree/master/docs/docs/02-developers/02-modules/reactions/06-permissions.md",tags:[],version:"current",sidebarPosition:6,frontMatter:{id:"permissions",title:"Permissions",sidebar_label:"Permissions",slug:"permissions"},sidebar:"docs",previous:{title:"Events",permalink:"/developers/modules/reactions/events"},next:{title:"Client",permalink:"/developers/modules/reactions/client"}},p={},c=[],m={toc:c};function d(e){var{components:t}=e,r=s(e,["components"]);return(0,n.kt)("wrapper",o({},m,r,{components:t,mdxType:"MDXLayout"}),(0,n.kt)("h1",o({},{id:"permissions"}),"Permissions"),(0,n.kt)("p",null,"Based on which ",(0,n.kt)("inlineCode",{parentName:"p"},"x/reactions")," related action your users want to perform, they need to be granted with one or more\nof the following permissions."),(0,n.kt)("table",null,(0,n.kt)("thead",{parentName:"table"},(0,n.kt)("tr",{parentName:"thead"},(0,n.kt)("th",o({parentName:"tr"},{align:"left"}),(0,n.kt)("strong",{parentName:"th"},"Permission Value")),(0,n.kt)("th",o({parentName:"tr"},{align:"left"}),(0,n.kt)("strong",{parentName:"th"},"Permission Description")))),(0,n.kt)("tbody",{parentName:"table"},(0,n.kt)("tr",{parentName:"tbody"},(0,n.kt)("td",o({parentName:"tr"},{align:"left"}),(0,n.kt)("inlineCode",{parentName:"td"},"ADD_REACTION")),(0,n.kt)("td",o({parentName:"tr"},{align:"left"}),"Allows to add and remove posts reactions")),(0,n.kt)("tr",{parentName:"tbody"},(0,n.kt)("td",o({parentName:"tr"},{align:"left"}),(0,n.kt)("inlineCode",{parentName:"td"},"MANAGE_REGISTERED_REACTIONS")),(0,n.kt)("td",o({parentName:"tr"},{align:"left"}),"Allows to manage the registered reactions")),(0,n.kt)("tr",{parentName:"tbody"},(0,n.kt)("td",o({parentName:"tr"},{align:"left"}),(0,n.kt)("inlineCode",{parentName:"td"},"MANAGE_REACTION_PARAMS")),(0,n.kt)("td",o({parentName:"tr"},{align:"left"}),"Allows to manage reaction parameters")))))}d.isMDXComponent=!0}}]);