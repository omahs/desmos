"use strict";(self.webpackChunkdesmos_docs=self.webpackChunkdesmos_docs||[]).push([[62993],{3905:(e,t,n)=>{n.d(t,{Zo:()=>l,kt:()=>u});var r=n(67294);function a(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function o(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function i(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?o(Object(n),!0).forEach((function(t){a(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):o(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function s(e,t){if(null==e)return{};var n,r,a=function(e,t){if(null==e)return{};var n,r,a={},o=Object.keys(e);for(r=0;r<o.length;r++)n=o[r],t.indexOf(n)>=0||(a[n]=e[n]);return a}(e,t);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(r=0;r<o.length;r++)n=o[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(a[n]=e[n])}return a}var c=r.createContext({}),p=function(e){var t=r.useContext(c),n=t;return e&&(n="function"==typeof e?e(t):i(i({},t),e)),n},l=function(e){var t=p(e.components);return r.createElement(c.Provider,{value:t},e.children)},d={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},m=r.forwardRef((function(e,t){var n=e.components,a=e.mdxType,o=e.originalType,c=e.parentName,l=s(e,["components","mdxType","originalType","parentName"]),m=p(n),u=a,f=m["".concat(c,".").concat(u)]||m[u]||d[u]||o;return n?r.createElement(f,i(i({ref:t},l),{},{components:n})):r.createElement(f,i({ref:t},l))}));function u(e,t){var n=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var o=n.length,i=new Array(o);i[0]=m;var s={};for(var c in t)hasOwnProperty.call(t,c)&&(s[c]=t[c]);s.originalType=e,s.mdxType="string"==typeof e?e:a,i[1]=s;for(var p=2;p<o;p++)i[p]=n[p];return r.createElement.apply(null,i)}return r.createElement.apply(null,n)}m.displayName="MDXCreateElement"},99621:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>p,contentTitle:()=>s,default:()=>m,frontMatter:()=>i,metadata:()=>c,toc:()=>l});n(67294);var r=n(3905);function a(){return a=Object.assign||function(e){for(var t=1;t<arguments.length;t++){var n=arguments[t];for(var r in n)Object.prototype.hasOwnProperty.call(n,r)&&(e[r]=n[r])}return e},a.apply(this,arguments)}function o(e,t){if(null==e)return{};var n,r,a=function(e,t){if(null==e)return{};var n,r,a={},o=Object.keys(e);for(r=0;r<o.length;r++)n=o[r],t.indexOf(n)>=0||(a[n]=e[n]);return a}(e,t);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(r=0;r<o.length;r++)n=o[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(a[n]=e[n])}return a}const i={id:"post-attachment",title:"post_attachment",hide_table_of_contents:!1},s=void 0,c={unversionedId:"graphql/objects/post-attachment",id:"version-4.2.0/graphql/objects/post-attachment",title:"post_attachment",description:'columns and relationships of "post_attachment"',source:"@site/versioned_docs/version-4.2.0/07-graphql/objects/post-attachment.mdx",sourceDirName:"07-graphql/objects",slug:"/graphql/objects/post-attachment",permalink:"/4.2.0/graphql/objects/post-attachment",draft:!1,editUrl:"https://github.com/desmos-labs/desmos/tree/master/docs/versioned_docs/version-4.2.0/07-graphql/objects/post-attachment.mdx",tags:[],version:"4.2.0",frontMatter:{id:"post-attachment",title:"post_attachment",hide_table_of_contents:!1},sidebar:"docs",previous:{title:"post_attachment_variance_fields",permalink:"/4.2.0/graphql/objects/post-attachment-variance-fields"},next:{title:"post_avg_fields",permalink:"/4.2.0/graphql/objects/post-avg-fields"}},p={},l=[{value:"Fields",id:"fields",level:3},{value:"<code>content</code> (<code>jsonb!</code>)",id:"content-jsonb",level:4},{value:"<code>id</code> (<code>bigint!</code>)",id:"id-bigint",level:4},{value:"<code>post</code> (<code>post!</code>)",id:"post-post",level:4}],d={toc:l};function m(e){var{components:t}=e,n=o(e,["components"]);return(0,r.kt)("wrapper",a({},d,n,{components:t,mdxType:"MDXLayout"}),(0,r.kt)("p",null,'columns and relationships of "post_attachment"'),(0,r.kt)("pre",null,(0,r.kt)("code",a({parentName:"pre"},{className:"language-graphql"}),"type post_attachment {\n  content(\n  path: String\n): jsonb!\n  id: bigint!\n  post: post!\n}\n")),(0,r.kt)("h3",a({},{id:"fields"}),"Fields"),(0,r.kt)("h4",a({},{id:"content-jsonb"}),(0,r.kt)("a",a({parentName:"h4"},{href:"#"}),(0,r.kt)("inlineCode",{parentName:"a"},"content"))," (",(0,r.kt)("a",a({parentName:"h4"},{href:"../scalars/jsonb"}),(0,r.kt)("inlineCode",{parentName:"a"},"jsonb!")),")"),(0,r.kt)("ul",null,(0,r.kt)("li",{parentName:"ul"},(0,r.kt)("h5",a({parentName:"li"},{id:"path-string"}),(0,r.kt)("a",a({parentName:"h5"},{href:"#"}),(0,r.kt)("inlineCode",{parentName:"a"},"path"))," (",(0,r.kt)("a",a({parentName:"h5"},{href:"../scalars/string"}),(0,r.kt)("inlineCode",{parentName:"a"},"String")),")"))),(0,r.kt)("p",null,"JSON select path"),(0,r.kt)("h4",a({},{id:"id-bigint"}),(0,r.kt)("a",a({parentName:"h4"},{href:"#"}),(0,r.kt)("inlineCode",{parentName:"a"},"id"))," (",(0,r.kt)("a",a({parentName:"h4"},{href:"../scalars/bigint"}),(0,r.kt)("inlineCode",{parentName:"a"},"bigint!")),")"),(0,r.kt)("h4",a({},{id:"post-post"}),(0,r.kt)("a",a({parentName:"h4"},{href:"#"}),(0,r.kt)("inlineCode",{parentName:"a"},"post"))," (",(0,r.kt)("a",a({parentName:"h4"},{href:"../objects/post"}),(0,r.kt)("inlineCode",{parentName:"a"},"post!")),")"),(0,r.kt)("p",null,"An object relationship"))}m.isMDXComponent=!0}}]);