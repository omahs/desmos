"use strict";(self.webpackChunkdesmos_docs=self.webpackChunkdesmos_docs||[]).push([[74651],{3905:(e,n,t)=>{t.d(n,{Zo:()=>u,kt:()=>g});var a=t(67294);function r(e,n,t){return n in e?Object.defineProperty(e,n,{value:t,enumerable:!0,configurable:!0,writable:!0}):e[n]=t,e}function o(e,n){var t=Object.keys(e);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);n&&(a=a.filter((function(n){return Object.getOwnPropertyDescriptor(e,n).enumerable}))),t.push.apply(t,a)}return t}function s(e){for(var n=1;n<arguments.length;n++){var t=null!=arguments[n]?arguments[n]:{};n%2?o(Object(t),!0).forEach((function(n){r(e,n,t[n])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(t)):o(Object(t)).forEach((function(n){Object.defineProperty(e,n,Object.getOwnPropertyDescriptor(t,n))}))}return e}function l(e,n){if(null==e)return{};var t,a,r=function(e,n){if(null==e)return{};var t,a,r={},o=Object.keys(e);for(a=0;a<o.length;a++)t=o[a],n.indexOf(t)>=0||(r[t]=e[t]);return r}(e,n);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(a=0;a<o.length;a++)t=o[a],n.indexOf(t)>=0||Object.prototype.propertyIsEnumerable.call(e,t)&&(r[t]=e[t])}return r}var i=a.createContext({}),c=function(e){var n=a.useContext(i),t=n;return e&&(t="function"==typeof e?e(n):s(s({},n),e)),t},u=function(e){var n=c(e.components);return a.createElement(i.Provider,{value:n},e.children)},p={inlineCode:"code",wrapper:function(e){var n=e.children;return a.createElement(a.Fragment,{},n)}},d=a.forwardRef((function(e,n){var t=e.components,r=e.mdxType,o=e.originalType,i=e.parentName,u=l(e,["components","mdxType","originalType","parentName"]),d=c(t),g=r,h=d["".concat(i,".").concat(g)]||d[g]||p[g]||o;return t?a.createElement(h,s(s({ref:n},u),{},{components:t})):a.createElement(h,s({ref:n},u))}));function g(e,n){var t=arguments,r=n&&n.mdxType;if("string"==typeof e||r){var o=t.length,s=new Array(o);s[0]=d;var l={};for(var i in n)hasOwnProperty.call(n,i)&&(l[i]=n[i]);l.originalType=e,l.mdxType="string"==typeof e?e:r,s[1]=l;for(var c=2;c<o;c++)s[c]=t[c];return a.createElement.apply(null,s)}return a.createElement.apply(null,t)}d.displayName="MDXCreateElement"},70964:(e,n,t)=>{t.r(n),t.d(n,{assets:()=>c,contentTitle:()=>l,default:()=>d,frontMatter:()=>s,metadata:()=>i,toc:()=>u});t(67294);var a=t(3905);function r(){return r=Object.assign||function(e){for(var n=1;n<arguments.length;n++){var t=arguments[n];for(var a in t)Object.prototype.hasOwnProperty.call(t,a)&&(e[a]=t[a])}return e},r.apply(this,arguments)}function o(e,n){if(null==e)return{};var t,a,r=function(e,n){if(null==e)return{};var t,a,r={},o=Object.keys(e);for(a=0;a<o.length;a++)t=o[a],n.indexOf(t)>=0||(r[t]=e[t]);return r}(e,n);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(a=0;a<o.length;a++)t=o[a],n.indexOf(t)>=0||Object.prototype.propertyIsEnumerable.call(e,t)&&(r[t]=e[t])}return r}const s={},l="ADR 017: Subspace fee grant",i={unversionedId:"architecture/adr-017-subspace-fee-grant",id:"architecture/adr-017-subspace-fee-grant",title:"ADR 017: Subspace fee grant",description:"Changelog",source:"@site/docs/architecture/adr-017-subspace-fee-grant.md",sourceDirName:"architecture",slug:"/architecture/adr-017-subspace-fee-grant",permalink:"/architecture/adr-017-subspace-fee-grant",draft:!1,editUrl:"https://github.com/desmos-labs/desmos/tree/master/docs/docs/architecture/adr-017-subspace-fee-grant.md",tags:[],version:"current",frontMatter:{},sidebar:"docs",previous:{title:"ADR 016: Reactions module",permalink:"/architecture/adr-016-reactions-module"},next:{title:"ADR 018: Improve subspace treasury account",permalink:"/architecture/adr-018-improve-subspace-treasury-account"}},c={},u=[{value:"Changelog",id:"changelog",level:2},{value:"Status",id:"status",level:2},{value:"Abstract",id:"abstract",level:2},{value:"Context",id:"context",level:2},{value:"Decision",id:"decision",level:2},{value:"DeductFeeDecorator",id:"deductfeedecorator",level:3},{value:"Types",id:"types",level:3},{value:"Allowance",id:"allowance",level:4},{value:"User Grant",id:"user-grant",level:4},{value:"Group Grant",id:"group-grant",level:4},{value:"Store",id:"store",level:3},{value:"Group grant",id:"group-grant-1",level:4},{value:"User grant",id:"user-grant-1",level:4},{value:"Expired Allowance",id:"expired-allowance",level:3},{value:"Store",id:"store-1",level:4},{value:"<code>Msg</code> Service",id:"msg-service",level:3},{value:"<code>Query</code> Service",id:"query-service",level:3},{value:"Consequences",id:"consequences",level:2},{value:"Backwards Compatibility",id:"backwards-compatibility",level:3},{value:"Positive",id:"positive",level:3},{value:"Negative",id:"negative",level:3},{value:"Neutral",id:"neutral",level:3},{value:"References",id:"references",level:2}],p={toc:u};function d(e){var{components:n}=e,t=o(e,["components"]);return(0,a.kt)("wrapper",r({},p,t,{components:n,mdxType:"MDXLayout"}),(0,a.kt)("h1",r({},{id:"adr-017-subspace-fee-grant"}),"ADR 017: Subspace fee grant"),(0,a.kt)("h2",r({},{id:"changelog"}),"Changelog"),(0,a.kt)("ul",null,(0,a.kt)("li",{parentName:"ul"},"Nov 14th, 2022: First draft;"),(0,a.kt)("li",{parentName:"ul"},"Nov 21th, 2022: First review.")),(0,a.kt)("h2",r({},{id:"status"}),"Status"),(0,a.kt)("p",null,"ACCEPTED Not Implemented"),(0,a.kt)("h2",r({},{id:"abstract"}),"Abstract"),(0,a.kt)("p",null,"This ADR introduces a new subspace-based fee grant method, which allows subspace owners/admins to pay the fees of subspace-related transactions on behalf of the users."),(0,a.kt)("h2",r({},{id:"context"}),"Context"),(0,a.kt)("p",null,"One of the major problems of current Web3-based applications that make them complicated to use by common people is the fact that they require users to have some tokens in order to pay transaction fees to perform various operations. For instance, in order to create a post on any Desmos-based social network, users have to:"),(0,a.kt)("ol",null,(0,a.kt)("li",{parentName:"ol"},"understand what a ",(0,a.kt)("em",{parentName:"li"},"transaction")," is and why they need some DSM to broadcast it;"),(0,a.kt)("li",{parentName:"ol"},"get some DSM either via an on-ramp service or by swapping existing funds.")),(0,a.kt)("p",null,"The ",(0,a.kt)("inlineCode",{parentName:"p"},"x/feegrant")," module of the Cosmos SDK allows anyone to pay fees on behalf of other users, meaning that the latter can ideally perform any kind of transaction without even knowing the concept of fees. However, the ",(0,a.kt)("inlineCode",{parentName:"p"},"x/feegrant")," allowance is not subspace-specific: this means that subspace admins/owners might unexpectingly pay the fees for transactions that are related to other subspaces."),(0,a.kt)("h2",r({},{id:"decision"}),"Decision"),(0,a.kt)("p",null,"We will implement a subspace-specific fee grant process, based on  the ",(0,a.kt)("inlineCode",{parentName:"p"},"x/feegrant")," module, that allows subspace fees providers to pay fees for the users inside the specified subspace. This system will allow subspace owners/admins to issue a fee grant towards either a users group or a single user, so that they can later leverage this grant to execute subspace-related transactions without having to worry about paying fees."),(0,a.kt)("h3",r({},{id:"deductfeedecorator"}),"DeductFeeDecorator"),(0,a.kt)("p",null,"Currently, the ",(0,a.kt)("inlineCode",{parentName:"p"},"x/auth")," module provides a ",(0,a.kt)("inlineCode",{parentName:"p"},"DeductFeeDecorator")," that deducts the fees of a transaction from the signer/feepayer. We will build a new subspace-specific ",(0,a.kt)("inlineCode",{parentName:"p"},"DeductFeeDecorator")," based on the existing one that acts as follows:"),(0,a.kt)("ol",null,(0,a.kt)("li",{parentName:"ol"},"check if all the messages in the transaction are messages related to the same subspace;"),(0,a.kt)("li",{parentName:"ol"},"perform the following operation:",(0,a.kt)("ol",{parentName:"li"},(0,a.kt)("li",{parentName:"ol"},"if all the messages are all related to the same subspace and a fee grant exists for the fee payer inside the subspace, run the ",(0,a.kt)("inlineCode",{parentName:"li"},"x/subspaces")," ",(0,a.kt)("inlineCode",{parentName:"li"},"DeductFeeDecorator")),(0,a.kt)("li",{parentName:"ol"},"otherwise, run the ",(0,a.kt)("inlineCode",{parentName:"li"},"x/auth")," ",(0,a.kt)("inlineCode",{parentName:"li"},"DeductFeeDecorator"))))),(0,a.kt)("h3",r({},{id:"types"}),"Types"),(0,a.kt)("h4",r({},{id:"allowance"}),"Allowance"),(0,a.kt)("p",null,"The ",(0,a.kt)("inlineCode",{parentName:"p"},"x/feegrant")," module already provides an interface (",(0,a.kt)("inlineCode",{parentName:"p"},"FeeAllowanceI"),") to represent a generic allowance, along with a set of useful implementations like ",(0,a.kt)("inlineCode",{parentName:"p"},"BasicAllowance"),", ",(0,a.kt)("inlineCode",{parentName:"p"},"AllowedMsgAllowance")," and ",(0,a.kt)("inlineCode",{parentName:"p"},"PeriodicAllowance"),". Since we are not going to introduce new kinds of allowances, nor we are going to edit how an allowance is represented, we will reuse these types."),(0,a.kt)("h4",r({},{id:"user-grant"}),"User Grant"),(0,a.kt)("p",null,"The ",(0,a.kt)("inlineCode",{parentName:"p"},"x/feegrant")," module provides a ",(0,a.kt)("inlineCode",{parentName:"p"},"Grant")," object used to store the ",(0,a.kt)("inlineCode",{parentName:"p"},"granter"),", ",(0,a.kt)("inlineCode",{parentName:"p"},"grantee")," and what kind of ",(0,a.kt)("inlineCode",{parentName:"p"},"allowance")," is granted to a specific user. Since this is enough information for us as well, we will use the same object to store user-related fee grants."),(0,a.kt)("h4",r({},{id:"group-grant"}),"Group Grant"),(0,a.kt)("p",null,"In order to represent a fee grant granted to a user group, we will implement a ",(0,a.kt)("inlineCode",{parentName:"p"},"GroupGrant")," type. This will contain the ",(0,a.kt)("inlineCode",{parentName:"p"},"granter"),", ",(0,a.kt)("inlineCode",{parentName:"p"},"group_id")," and what kind of ",(0,a.kt)("inlineCode",{parentName:"p"},"allowance")," is granted to the group:"),(0,a.kt)("pre",null,(0,a.kt)("code",r({parentName:"pre"},{className:"language-protobuf"}),'message GroupGrant {\n    // granter is the address of the user granting an allowance of their funds.\n  string granter = 1 [(cosmos_proto.scalar) = "cosmos.AddressString"];\n\n  // the id of the group being granted an allowance of another user\'s funds.\n  uint32 group_id = 2;\n\n  // allowance can be any of basic, periodic, allowed fee allowance.\n  google.protobuf.Any allowance = 3 [(cosmos_proto.accepts_interface) = "FeeAllowanceI"];\n}\n')),(0,a.kt)("h3",r({},{id:"store"}),"Store"),(0,a.kt)("h4",r({},{id:"group-grant-1"}),"Group grant"),(0,a.kt)("p",null,"In order to simplify the management of group allowances, each granted group allowance will be stored using the following key:"),(0,a.kt)("pre",null,(0,a.kt)("code",r({parentName:"pre"},{}),"SubspaceGroupAllowancePrefix | SubspaceID | GranterAddress | GroupID | -> Protobuf(GroupGrant)\n")),(0,a.kt)("p",null,"This structure allows granters to easily manage the group allowance inside a subspace by iterating over all allowances for the granters, which will be the most used query. On the other hand, it allows to perform the following queries: "),(0,a.kt)("ol",null,(0,a.kt)("li",{parentName:"ol"},"finding a specific grant given the granter address.",(0,a.kt)("br",{parentName:"li"}),"This can be done in O(N), where N is the number of granted groups by the granter."),(0,a.kt)("li",{parentName:"ol"},"finding all the grants given to a user based on the groups they are part of.\nThis can be done in O(N * log(M)), where N is the number of all the grants inside the subspace, and M is the number of members in the user group.")),(0,a.kt)("p",null,"Considering that: "),(0,a.kt)("ul",null,(0,a.kt)("li",{parentName:"ul"},"only users with permissions will be able to create grants in each subspace"),(0,a.kt)("li",{parentName:"ul"},"there will be a small number of grants each granter will create")),(0,a.kt)("p",null,"then the time complexity of searching if a user has been granted a fee grant should not become too large."),(0,a.kt)("h4",r({},{id:"user-grant-1"}),"User grant"),(0,a.kt)("p",null,"Each subspace fee granted allowance will be stored using keys with the following structure:"),(0,a.kt)("pre",null,(0,a.kt)("code",r({parentName:"pre"},{}),"SubspaceUserAllowancePrefix | SubspaceID | GranterAddress | GranteeAddress |-> Protobuf(Grant)\n")),(0,a.kt)("p",null,"This structure allows granters to easily manage their grants inside a subspace by iterating over all grants for the granters, which will be the most used query. On the other hand, it allows the following queries:"),(0,a.kt)("ol",null,(0,a.kt)("li",{parentName:"ol"},"retrieving a specific grant.",(0,a.kt)("br",{parentName:"li"}),"This will require grantees to know who the granter address, and can be done in O(1)."),(0,a.kt)("li",{parentName:"ol"},"retrieving all the grants that a user has been granted.\nThis can be done by iterating over all the granters' grants in O(N), where N is the number of all the grants inside the given subspace.")),(0,a.kt)("h3",r({},{id:"expired-allowance"}),"Expired Allowance"),(0,a.kt)("p",null,"The subspace owners may want to set the allowance to expire after a certain amount of time. In order to reduce the database size, we will use a custom ",(0,a.kt)("inlineCode",{parentName:"p"},"BeginBlocker")," to remove all expired allowances."),(0,a.kt)("h4",r({},{id:"store-1"}),"Store"),(0,a.kt)("p",null,"In order to make it easier to remove expired allowances, we will implement queues sorted by expiration time. Then, during the ",(0,a.kt)("inlineCode",{parentName:"p"},"BeginBlock")," phase, we iterate over all expired keys before the block time and remove them. The structure of the queues keys will be the following:"),(0,a.kt)("pre",null,(0,a.kt)("code",r({parentName:"pre"},{}),"ExpiringAllowanceQueuePrefix | Expired Time | Subspace ID | Granter Address | (User Address or Group ID) -> 0x01\n")),(0,a.kt)("h3",r({},{id:"msg-service"}),(0,a.kt)("inlineCode",{parentName:"h3"},"Msg")," Service"),(0,a.kt)("p",null,"In order to allow subspace fees providers to grant an allowance for the users, we will have the following operations:"),(0,a.kt)("ul",null,(0,a.kt)("li",{parentName:"ul"},"grant an allowance to a user"),(0,a.kt)("li",{parentName:"ul"},"revoke an allowance to a user"),(0,a.kt)("li",{parentName:"ul"},"grant an allowance to a group"),(0,a.kt)("li",{parentName:"ul"},"revoke an allowance to a group")),(0,a.kt)("pre",null,(0,a.kt)("code",r({parentName:"pre"},{className:"language-protobuf"}),"service Msg {\n    // GrantUserAllowance allows the granter to grant a fee allowance to the grantee.\n    rpc GrantUserAllowance(MsgGrantUserAllowance) returns(MsgGrantUserAllowanceResponse);\n\n    // RevokeUserAllowance allows a granter to revoke any existing allowance that has to been granted to the grantee.\n    rpc RevokeUserAllowance(MsgRevokeUserAllowance) returns(MsgRevokeUserAllowanceResponse);\n\n    // GrantGroupAllowance allows the granter to grant a fee allowance to the group.\n    rpc GrantGroupAllowance(MsgGrantGroupAllowance) returns(MsgGrantGroupAllowanceResponse);\n\n    // RevokeGroupAllowance allows a granter to revoke any existing allowance that has to been granted to the group.\n    rpc RevokeGroupAllowance(MsgRevokeGroupAllowance) returns(MsgRevokeGroupAllowanceResponse);\n}\n\n// MsgGrantUserAllowance adds permissions for the grantee to spend up allowance of fees from the granter inside the given subspace.\nmessage MsgGrantUserAllowance {\n    // the id of the subspace where the granter grants the allowance to the grantee.\n    uint64 subspace_id = 1;\n\n    // the address of the user granting an allowance of their funds.\n    string granter = 2;\n\n    // the address of the user being granted an allowance of another user's funds.\n    string grantee = 3;\n\n    // allowance can be any of fee allowances which implements FeeAllowanceI.\n    google.protobuf.Any allowance = 4 [(cosmos_proto.accepts_interface) = \"FeeAllowanceI\"];\n}\n\n// MsgGrantUserAllowanceResponse defines the Msg/GrantAllowanceResponse response type.\nmessage MsgGrantUserAllowanceResponse {}\n\n// MsgRevokeUserAllowance removes any existing allowance from granter to the grantee inside the subspace.\nmessage MsgRevokeUserAllowance {\n    // the id of the subspace where the granter grants the allowance to the grantee.\n    uint64 subspace_id = 1;\n\n    // the address of the user granting an allowance of their funds.\n    string granter = 2;\n\n    // the address of the user being granted an allowance of another user's funds.\n    string grantee = 3;\n}\n\n// MsgRevokeUserAllowanceResponse defines the Msg/RevokeAllowanceResponse response type.\nmessage MsgRevokeUserAllowanceResponse {}\n\n// MsgGrantGroupAllowance adds permissions for the group to spend up allowance of fees from the granter inside the given subspace.\nmessage MsgGrantGroupAllowance {\n    // the id of the subspace where the granter grants the allowance to the grantee.\n    uint64 subspace_id = 1;\n\n    // the id of the group being granted an allowance of another user's funds.  \n    uint32 group_id = 2;\n\n    // the address of the user granting an allowance of their funds.\n    string granter = 3;\n\n    // allowance can be any of fee allowances which implements FeeAllowanceI.\n    google.protobuf.Any allowance = 4 [(cosmos_proto.accepts_interface) = \"FeeAllowanceI\"];\n}\n\n// MsgGrantGroupAllowanceResponse defines the Msg/GrantAllowanceResponse response type.\nmessage MsgGrantGroupAllowanceResponse {}\n\n// MsgRevokeGroupAllowance removes any existing allowance from granter to the group inside the subspace.\nmessage MsgRevokeGroupAllowance {\n    // the id of the subspace where the granter grants the allowance to the group.\n    uint64 subspace_id = 1;\n\n    // the id of the group being granted an allowance of another user's funds.\n    uint32 group_id = 2;\n\n    // the address of the user granting an allowance of their funds.\n    string granter = 3;\n}\n\n// MsgRevokeGroupAllowanceResponse defines the Msg/RevokeAllowanceResponse response type.\nmessage MsgRevokeUserAllowanceResponse {}\n")),(0,a.kt)("h3",r({},{id:"query-service"}),(0,a.kt)("inlineCode",{parentName:"h3"},"Query")," Service"),(0,a.kt)("p",null,"In order to allow clients to easily query for allowances we will implement the following new queries:"),(0,a.kt)("pre",null,(0,a.kt)("code",r({parentName:"pre"},{className:"language-protobuf"}),"service Query {\n    // UserAllowances returns all the grants for users.\n    rpc UserAllowances(QueryUserAllowancesRequest) returns (QueryAllowancesResponse) {\n        option (google.api.http).get = \"/desmos/subspaces/v3/subspaces/{subspace_id}/allowances/users\";\n    }\n    // GroupAllowances returns all the grants for groups.\n    rpc GroupAllowances(QueryGroupAllowancesRequest) returns(QueryGroupAllowancesResponse) {\n        option (google.api.http).get = \"/desmos/subspaces/v3/subspaces/{subspace_id}/allowances/groups\";\n    }\n}\n\n// QueryUserAllowancesRequest is the request type for the Query/UserAllowances RPC method.\nmessage QueryUserAllowancesRequest {\n    // the id of the subspace where the granter grants the allowance to the grantee.\n    uint64 subspace_id = 1;\n\n    // (optional) the address of the user granting an allowance of their funds.\n    string granter = 2;\n\n    // (Optional) the address of the user being granted an allowance of another user's funds.\n    string grantee = 3;\n\n    // pagination defines an pagination for the request.\n    cosmos.base.query.v1beta1.PageRequest pagination = 4;\n}\n\n// QueryUserAllowancesResponse is the response type for the Query/UserAllowances RPC method.\nmessage QueryUserAllowancesResponse {\n    // allowances are allowance's granted for grantee by granter.\n    repeated cosmos.feegrant.v1beta1.Grant allowances = 1;\n\n    // pagination defines an pagination for the response.\n    cosmos.base.query.v1beta1.PageResponse pagination = 2;\n}\n\n// QueryGroupAllowancesRequest is the request type for the Query/GroupAllowances RPC method.\nmessage QueryGroupAllowancesRequest {\n    // the id of the subspace where the granter grants the allowance to the grantee.\n    uint64 subspace_id = 1;\n\n    // (optional) the address of the user granting an allowance of their funds.\n    string granter = 2;\n\n    // (Optional) the id of the group being granted an allowance of another user's funds.\n    uint32 group_id = 3;\n\n    // pagination defines an pagination for the request.\n    cosmos.base.query.v1beta1.PageRequest pagination = 4;\n}\n\n// QueryGroupAllowancesResponse is the response type for the Query/GroupAllowances RPC method.\nmessage QueryGroupAllowancesResponse {\n    // allowances are allowance's granted for grantee by granter.\n    repeated cosmos.subspace.v3.GroupGrant allowances = 1;\n\n    // pagination defines an pagination for the response.\n    cosmos.base.query.v1beta1.PageResponse pagination = 2;\n}\n")),(0,a.kt)("h2",r({},{id:"consequences"}),"Consequences"),(0,a.kt)("h3",r({},{id:"backwards-compatibility"}),"Backwards Compatibility"),(0,a.kt)("p",null,"The proposed solution introduces a new set of store keys, thus it is completely backward compatible."),(0,a.kt)("h3",r({},{id:"positive"}),"Positive"),(0,a.kt)("ul",null,(0,a.kt)("li",{parentName:"ul"},"Allow subspace fee providers to grant fees allowances to their users.")),(0,a.kt)("h3",r({},{id:"negative"}),"Negative"),(0,a.kt)("ul",null,(0,a.kt)("li",{parentName:"ul"},"Storing extra subspace grant info takes up more storage space;"),(0,a.kt)("li",{parentName:"ul"},"Having more checks during the fee deduction phase slows down the transaction processing operation;"),(0,a.kt)("li",{parentName:"ul"},"Removing expired allowances at the block begin phase increases the block production time.  ")),(0,a.kt)("h3",r({},{id:"neutral"}),"Neutral"),(0,a.kt)("ul",null,(0,a.kt)("li",{parentName:"ul"},"Not known")),(0,a.kt)("h2",r({},{id:"references"}),"References"),(0,a.kt)("ul",null,(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("a",r({parentName:"li"},{href:"https://docs.cosmos.network/v0.46/modules/feegrant/01_concepts.html"}),"Cosmos SDK fee grant concepts"))))}d.isMDXComponent=!0}}]);