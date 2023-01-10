"use strict";(self.webpackChunkdesmos_docs=self.webpackChunkdesmos_docs||[]).push([[16394],{3905:(e,t,n)=>{n.d(t,{Zo:()=>p,kt:()=>d});var r=n(67294);function i(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function o(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function s(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?o(Object(n),!0).forEach((function(t){i(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):o(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function a(e,t){if(null==e)return{};var n,r,i=function(e,t){if(null==e)return{};var n,r,i={},o=Object.keys(e);for(r=0;r<o.length;r++)n=o[r],t.indexOf(n)>=0||(i[n]=e[n]);return i}(e,t);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(r=0;r<o.length;r++)n=o[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(i[n]=e[n])}return i}var l=r.createContext({}),c=function(e){var t=r.useContext(l),n=t;return e&&(n="function"==typeof e?e(t):s(s({},t),e)),n},p=function(e){var t=c(e.components);return r.createElement(l.Provider,{value:t},e.children)},u={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},f=r.forwardRef((function(e,t){var n=e.components,i=e.mdxType,o=e.originalType,l=e.parentName,p=a(e,["components","mdxType","originalType","parentName"]),f=c(n),d=i,m=f["".concat(l,".").concat(d)]||f[d]||u[d]||o;return n?r.createElement(m,s(s({ref:t},p),{},{components:n})):r.createElement(m,s({ref:t},p))}));function d(e,t){var n=arguments,i=t&&t.mdxType;if("string"==typeof e||i){var o=n.length,s=new Array(o);s[0]=f;var a={};for(var l in t)hasOwnProperty.call(t,l)&&(a[l]=t[l]);a.originalType=e,a.mdxType="string"==typeof e?e:i,s[1]=a;for(var c=2;c<o;c++)s[c]=n[c];return r.createElement.apply(null,s)}return r.createElement.apply(null,n)}f.displayName="MDXCreateElement"},23759:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>c,contentTitle:()=>a,default:()=>f,frontMatter:()=>s,metadata:()=>l,toc:()=>p});n(67294);var r=n(3905);function i(){return i=Object.assign||function(e){for(var t=1;t<arguments.length;t++){var n=arguments[t];for(var r in n)Object.prototype.hasOwnProperty.call(n,r)&&(e[r]=n[r])}return e},i.apply(this,arguments)}function o(e,t){if(null==e)return{};var n,r,i=function(e,t){if(null==e)return{};var n,r,i={},o=Object.keys(e);for(r=0;r<o.length;r++)n=o[r],t.indexOf(n)>=0||(i[n]=e[n]);return i}(e,t);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(r=0;r<o.length;r++)n=o[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(i[n]=e[n])}return i}const s={id:"genesis-file",title:"Genesis File",sidebar_label:"Genesis File",slug:"genesis-file"},a="Genesis file",l={unversionedId:"mainnet/genesis-file",id:"mainnet/genesis-file",title:"Genesis File",description:"If you are looking for testnet genesis file, please refer to",source:"@site/docs/06-mainnet/01-genesis-file.md",sourceDirName:"06-mainnet",slug:"/mainnet/genesis-file",permalink:"/mainnet/genesis-file",draft:!1,editUrl:"https://github.com/desmos-labs/desmos/tree/master/docs/docs/06-mainnet/01-genesis-file.md",tags:[],version:"current",lastUpdatedAt:1670002176,formattedLastUpdatedAt:"Dec 2, 2022",sidebarPosition:1,frontMatter:{id:"genesis-file",title:"Genesis File",sidebar_label:"Genesis File",slug:"genesis-file"},sidebar:"docs",previous:{title:"Full Sync",permalink:"/testnet/join-public/full-sync"},next:{title:"Seeds",permalink:"/mainnet/seeds"}},c={},p=[{value:"1. Download",id:"1-download",level:2},{value:"2. Check",id:"2-check",level:2}],u={toc:p};function f(e){var{components:t}=e,n=o(e,["components"]);return(0,r.kt)("wrapper",i({},u,n,{components:t,mdxType:"MDXLayout"}),(0,r.kt)("h1",i({},{id:"genesis-file"}),"Genesis file"),(0,r.kt)("admonition",i({},{title:"Mainnet only   ",type:"caution"}),(0,r.kt)("p",{parentName:"admonition"},"If you are looking for testnet genesis file, please refer to\n",(0,r.kt)("a",i({parentName:"p"},{href:"/testnet/join-public/genesis-file"}),"this")," instead.  ")),(0,r.kt)("p",null,"To connect to the ",(0,r.kt)("inlineCode",{parentName:"p"},"desmos-mainnet"),", you will need to download the correct genesis file."),(0,r.kt)("h2",i({},{id:"1-download"}),"1. Download"),(0,r.kt)("p",null,"Visit the ",(0,r.kt)("a",i({parentName:"p"},{href:"https://github.com/desmos-labs/mainnet"}),"mainnet repo")," and\ndownload the correct genesis file by running the following command."),(0,r.kt)("pre",null,(0,r.kt)("code",i({parentName:"pre"},{className:"language-bash"}),"# Download the existing genesis file for the mainnet\n# Replace <chain-id> with the id of the testnet you would like to join\ncurl https://raw.githubusercontent.com/desmos-labs/mainnet/main/genesis.json > ~/.desmos/config/genesis.json\n")),(0,r.kt)("h2",i({},{id:"2-check"}),"2. Check"),(0,r.kt)("p",null,"After the download, ensure it's the correct one by checking that it has the correct hashsum like below:"),(0,r.kt)("pre",null,(0,r.kt)("code",i({parentName:"pre"},{className:"language-bash"}),"jq -S -c -M '' ~/.desmos/config/genesis.json | shasum -a 256\n")))}f.isMDXComponent=!0}}]);