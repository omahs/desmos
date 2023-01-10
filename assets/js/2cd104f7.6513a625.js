"use strict";(self.webpackChunkdesmos_docs=self.webpackChunkdesmos_docs||[]).push([[97526],{3905:(e,t,r)=>{r.d(t,{Zo:()=>u,kt:()=>h});var n=r(67294);function s(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function o(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,n)}return r}function a(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?o(Object(r),!0).forEach((function(t){s(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):o(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}function i(e,t){if(null==e)return{};var r,n,s=function(e,t){if(null==e)return{};var r,n,s={},o=Object.keys(e);for(n=0;n<o.length;n++)r=o[n],t.indexOf(r)>=0||(s[r]=e[r]);return s}(e,t);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(n=0;n<o.length;n++)r=o[n],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(s[r]=e[r])}return s}var p=n.createContext({}),l=function(e){var t=n.useContext(p),r=t;return e&&(r="function"==typeof e?e(t):a(a({},t),e)),r},u=function(e){var t=l(e.components);return n.createElement(p.Provider,{value:t},e.children)},d={inlineCode:"code",wrapper:function(e){var t=e.children;return n.createElement(n.Fragment,{},t)}},c=n.forwardRef((function(e,t){var r=e.components,s=e.mdxType,o=e.originalType,p=e.parentName,u=i(e,["components","mdxType","originalType","parentName"]),c=l(r),h=s,m=c["".concat(p,".").concat(h)]||c[h]||d[h]||o;return r?n.createElement(m,a(a({ref:t},u),{},{components:r})):n.createElement(m,a({ref:t},u))}));function h(e,t){var r=arguments,s=t&&t.mdxType;if("string"==typeof e||s){var o=r.length,a=new Array(o);a[0]=c;var i={};for(var p in t)hasOwnProperty.call(t,p)&&(i[p]=t[p]);i.originalType=e,i.mdxType="string"==typeof e?e:s,a[1]=i;for(var l=2;l<o;l++)a[l]=r[l];return n.createElement.apply(null,a)}return n.createElement.apply(null,r)}c.displayName="MDXCreateElement"},5795:(e,t,r)=>{r.r(t),r.d(t,{assets:()=>l,contentTitle:()=>i,default:()=>c,frontMatter:()=>a,metadata:()=>p,toc:()=>u});r(67294);var n=r(3905);function s(){return s=Object.assign||function(e){for(var t=1;t<arguments.length;t++){var r=arguments[t];for(var n in r)Object.prototype.hasOwnProperty.call(r,n)&&(e[n]=r[n])}return e},s.apply(this,arguments)}function o(e,t){if(null==e)return{};var r,n,s=function(e,t){if(null==e)return{};var r,n,s={},o=Object.keys(e);for(n=0;n<o.length;n++)r=o[n],t.indexOf(r)>=0||(s[r]=e[r]);return s}(e,t);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(n=0;n<o.length;n++)r=o[n],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(s[r]=e[r])}return s}const a={},i="ADR 011: Reports module",p={unversionedId:"architecture/adr-011-reports-module",id:"architecture/adr-011-reports-module",title:"ADR 011: Reports module",description:"Changelog",source:"@site/docs/architecture/adr-011-reports-module.md",sourceDirName:"architecture",slug:"/architecture/adr-011-reports-module",permalink:"/architecture/adr-011-reports-module",draft:!1,editUrl:"https://github.com/desmos-labs/desmos/tree/master/docs/docs/architecture/adr-011-reports-module.md",tags:[],version:"current",frontMatter:{},sidebar:"docs",previous:{title:"ADR 010: Posts module",permalink:"/architecture/adr-010-posts-module"},next:{title:"ADR 012: Subspaces sections",permalink:"/architecture/adr-012-subspaces-sections"}},l={},u=[{value:"Changelog",id:"changelog",level:2},{value:"Status",id:"status",level:2},{value:"Abstract",id:"abstract",level:2},{value:"Context",id:"context",level:2},{value:"Decision",id:"decision",level:2},{value:"Types",id:"types",level:3},{value:"Report",id:"report",level:4},{value:"Reason",id:"reason",level:3},{value:"Params",id:"params",level:3},{value:"<code>Msg</code> Service",id:"msg-service",level:3},{value:"<code>Query</code> Service",id:"query-service",level:3},{value:"Consequences",id:"consequences",level:2},{value:"Backwards Compatibility",id:"backwards-compatibility",level:3},{value:"Positive",id:"positive",level:3},{value:"Negative",id:"negative",level:3},{value:"Neutral",id:"neutral",level:3},{value:"Further Discussions",id:"further-discussions",level:2},{value:"References",id:"references",level:2}],d={toc:u};function c(e){var{components:t}=e,r=o(e,["components"]);return(0,n.kt)("wrapper",s({},d,r,{components:t,mdxType:"MDXLayout"}),(0,n.kt)("h1",s({},{id:"adr-011-reports-module"}),"ADR 011: Reports module"),(0,n.kt)("h2",s({},{id:"changelog"}),"Changelog"),(0,n.kt)("ul",null,(0,n.kt)("li",{parentName:"ul"},"April 12th, 2022: Initial draft")),(0,n.kt)("h2",s({},{id:"status"}),"Status"),(0,n.kt)("p",null,"ACCEPTED Implemented"),(0,n.kt)("h2",s({},{id:"abstract"}),"Abstract"),(0,n.kt)("p",null,"This ADR contains the specification of the ",(0,n.kt)("inlineCode",{parentName:"p"},"x/reports")," module which will allow users to report other users and/or content posted inside subspaces."),(0,n.kt)("h2",s({},{id:"context"}),"Context"),(0,n.kt)("p",null,"One of the most important parts of a functional social network is the set of rules that constitute the Terms of Service, and how they are enforced. In most social networks the key for a good user experience is the reporting system: users are able to report misbehaving users or inappropriate content that gets later reviewed and taken care of by moderators. "),(0,n.kt)("p",null,"Inside Desmos we MUST give users a set of tools that allow them to perform the same actions on all the social networks that they will have access to. At the same time, we SHOULD NOT enforce any high-level Term of Service since each subspace that will be built on Desmos might want to have very different rules from the other ones (i.e. an adult content social network will most likely have a very different set of rules from a kids social network). Instead, we should allow subspace owners and admins to register the various reasons a user/content can be reported for."),(0,n.kt)("h2",s({},{id:"decision"}),"Decision"),(0,n.kt)("p",null,"We will implement a new module named ",(0,n.kt)("inlineCode",{parentName:"p"},"x/reports")," that allows users to report either a misbehaving user or a bad content inside subspaces where they have the permission to do so."),(0,n.kt)("p",null,"The same module will also allow subspace owners to register their own custom supported reasons for reports. "),(0,n.kt)("h3",s({},{id:"types"}),"Types"),(0,n.kt)("p",null,"Reports can be of two different types: ",(0,n.kt)("inlineCode",{parentName:"p"},"UserReport")," for reporting a misbehaving user, or ",(0,n.kt)("inlineCode",{parentName:"p"},"PostReport")," for reporting a bad post. The whole system must be thought in order to support additional report types in the future (i.e. ",(0,n.kt)("inlineCode",{parentName:"p"},"ProposalReport")," to report a bad on-chain governance proposal). "),(0,n.kt)("h4",s({},{id:"report"}),"Report"),(0,n.kt)("pre",null,(0,n.kt)("code",s({parentName:"pre"},{className:"language-protobuf"}),"// Report contains the data of a generic report\nmessage Report {\n  // Id of the report inside the subspace\n  required uint64 id = 1;\n  \n  // Id of the reason this report has been created for\n  required uint32 reason_id = 2;\n  \n  // Message attached to this report\n  optional string message = 3;\n  \n  // Address of the reporter \n  required string reporter = 4;\n\n  // Target of the report\n  oneof Target {\n    UserTarget user_data = 5;\n    PostTarget post_data = 6;\n  }\n}\n\n// UserTarget contains the data of a report about a user\nmessage UserTarget {\n  // Id of the subspace inside which the user has been reported\n  required uint64 subspace_id = 1;\n  \n  // Address of the reported user\n  required string user = 2;\n}\n\n// PostTarget contains the data of a report about a post\nmessage PostTarget {\n  // Id of the subspace inside which the reported post is\n  required uint64 subspace_id = 1;\n  \n  // Id of the reported post \n  required uint64 post_id = 2;\n}\n")),(0,n.kt)("h3",s({},{id:"reason"}),"Reason"),(0,n.kt)("p",null,"Each subspace owner/admin SHOULD be able to define a custom set of supported reasons for which users/posts can be reported for. Each reason MUST contain a human-readable title that identifies the reason itself. "),(0,n.kt)("pre",null,(0,n.kt)("code",s({parentName:"pre"},{className:"language-protobuf"}),"// Reason contains the data about a reporting reason\nmessage Reason {\n  // Id of the reason inside the subspace \n  required uint32 id = 1;\n  \n  // Title of the reason\n  required string title = 2;\n  \n  // Extended description of the reason and the cases it applies to\n  optional string description = 3;\n}\n")),(0,n.kt)("h3",s({},{id:"params"}),"Params"),(0,n.kt)("p",null,"In order to make it easier to set up initially supported reporting reasons for new subspaces, we are going to provide all the subspaces with a list from which they can easily pick reasons that should be supported inside their own subspace as well."),(0,n.kt)("pre",null,(0,n.kt)("code",s({parentName:"pre"},{className:"language-protobuf"}),"// Params contains the module parameters\nmessage Params {\n  // List of available reasons from which new subspaces can pick their default ones\n  repeated Reason reasons = 1;\n}\n")),(0,n.kt)("h3",s({},{id:"msg-service"}),(0,n.kt)("inlineCode",{parentName:"h3"},"Msg")," Service"),(0,n.kt)("p",null,"We will allow the following operations: "),(0,n.kt)("ul",null,(0,n.kt)("li",{parentName:"ul"},"crete a new report"),(0,n.kt)("li",{parentName:"ul"},"delete an existing report "),(0,n.kt)("li",{parentName:"ul"},"manage supported reporting reasons (add, remove a reason)")),(0,n.kt)("blockquote",null,(0,n.kt)("p",{parentName:"blockquote"},"NOTE",(0,n.kt)("br",{parentName:"p"}),"\n","The ability to edit a report will ",(0,n.kt)("strong",{parentName:"p"},"not")," be allowed in order to avoid ever-changing reports that can make the moderation work a lot more complicated.")),(0,n.kt)("pre",null,(0,n.kt)("code",s({parentName:"pre"},{className:"language-protobuf"}),"service Msg {\n  // CreateReport allows to create a new report\n  rpc CreateReport(MsgCreateReport) returns (MsgCreateReportResponse);\n  \n  // DeleteReport allows to delete an existing report\n  rpc DeleteReport(MsgDeleteReport) returns (MsgDeleteReportResponse);\n  \n  // SupportReason allows to support one of the reasons present inside the module params \n  rpc SupportReason(MsgSupportReason) returns (MsgSupportReasonResponse);\n  \n  // AddReason allows to add a new supported reporting reason\n  rpc AddReason(MsgAddReason) returns (MsgAddReasonResponse);\n  \n  // RemoveReason allows to remove a supported reporting reason\n  rpc RemoveReason(MsgRemoveReason) returns (MsgRemoveReasonResponse);\n}\n\n// MsgCreateReport represents the message to be used to create a report\nmessage MsgCreateReport {\n  // Id of the reason this report has been created for\n  required uint32 reason_id = 1;\n\n  // Message attached to this report\n  optional string message = 2;\n\n  // Address of the reporter \n  required string reporter = 3;\n\n  // Target of the report\n  oneof Target {\n    UserTarget user_data = 4;\n    PostTarget post_data = 5;\n  }\n}\n\n// MsgCreateReportResponse represents the Msg/CreateReport response type\nmessage MsgCreateReportResponse {\n  // Id of the newly created report\n  required uint64 report_id = 1;\n}\n\n// MsgDeleteReport represents the message to be used when deleting a report\nmessage MsgDeleteReport {\n  // Id of the subspace that contains the report to be deleted \n  required uint64 subspace_id = 1;\n  \n  // Id of the report to be deleted \n  required uint64 report_id = 2; \n  \n  // Address of the user deleting the report \n  required string signer = 3;\n}\n\n// MsgDeleteReportResponse represents the Msg/DeleteReport response type\nmessage MsgDeleteReportResponse {}\n\n// MsgSupportReason represents the message to be used when wanting to support a reason from the module params\nmessage MsgSupportReason {\n  // Id of the subspace for which to support the reason\n  required uint64 subspace_id = 1;\n  \n  // Id of the reason that should be supported\n  required uint32 reason_id = 2;\n  \n  // Address of the user signing the message\n  required string signer = 3;\n}\n\n// MsgSupportReasonResponse represents the Msg/SupportReason response type\nmessage MsgSupportReasonResponse {}\n\n// MsgAddReason represents the message to be used when adding a new supported reason\nmessage MsgAddReason {\n  // Id of the subspace for which to add the reason \n  required uint64 subspace_id = 1;\n  \n  // Title of the reason\n  required string title = 2;\n\n  // Extended description of the reason and the cases it applies to\n  optional string description = 3;\n  \n  // Address of the user adding the supported reason\n  required string signer = 4;\n}\n\n// MsgAddReasonResponse represents the Msg/AddReason response type\nmessage MsgAddReasonResponse {\n  // Id of the newly supported reason\n  required uint32 reason_id = 1;\n}\n\n// MsgRemoveReason represents the message to be used when removing an exiting reporting reason\nmessage MsgRemoveReason {\n  // Id of the subspace from which to remove the reason \n  required uint64 subspace_id = 1;\n\n  // Id of the reason to be deleted\n  required uint32 reason_id = 2;\n\n  // Address of the user adding the supported reason\n  required string signer = 3;\n}\n\n// MsgRemoveReasonResponse represents the Msg/RemoveReason response type\nmessage MsgRemoveReasonResponse {\n  // Id of the newly supported reason\n  required uint32 reason_id = 1;\n}\n\n")),(0,n.kt)("h3",s({},{id:"query-service"}),(0,n.kt)("inlineCode",{parentName:"h3"},"Query")," Service"),(0,n.kt)("pre",null,(0,n.kt)("code",s({parentName:"pre"},{className:"language-protobuf"}),'service Query {\n  // Reports allows to query the reports for a specific target\n  rpc Reports(QueryReportsRequest) returns (QueryReportsResponse) {\n    option (google.api.http).get = "/desmos/reports/v1/{subspace_id}/reports";\n  }\n  \n  // Reasons allows to query the supported reporting reasons for a subspace\n  rpc Reasons(QueryReasonsRequest) returns (QueryReasonsResponse) {\n    option (google.api.http).get = "/desmos/reports/v1/{subspace_id}/reasons";\n  }\n  \n  // Params allows to query the module parameters\n  rpc Params(QueryParamsRequest) returns (QueryParamsResponse) {\n    option (google.api.http).get = "/desmos/reports/v1/params";\n  }\n}\n\n// QueryReportsResponse is the request type for Query/Reports RPC method\nmessage QueryReportsRequest {\n  // Target to query the reports for\n  oneof Target {\n    UserTarget user_data = 1;\n    PostTarget post_data = 2;\n  }\n\n  // pagination defines an optional pagination for the request.\n  optional cosmos.base.query.v1beta1.PageRequest pagination = 3;\n} \n\n// QueryReportsResponse is the response type for Query/Reports RPC method\nmessage QueryReportsResponse {\n  repeated Report reports = 1;\n  required cosmos.base.query.v1beta1.PageResponse pagination = 2;\n}\n\n// QueryReasonsRequest is the request type for Query/Reasons RPC method\nmessage QueryReasonsRequest {\n  // Id of the subspace to query the supported reporting reasons for\n  required uint64 subspace_id = 1;\n  \n  // pagination defines an optional pagination for the request.\n  optional cosmos.base.query.v1beta1.PageRequest pagination = 3;\n}\n\n// QueryReasonsResponse is the response type for Query/Reasons RPC method\nmessage QueryReasonsResponse {\n  repeated Reason reasons = 1;\n  required cosmos.base.query.v1beta1.PageResponse pagination = 2;\n}\n\n// QueryParamsRequest is the request type for Query/Params RPC method\nmessage QueryParamsRequest {}\n\n// QueryParamsResponse is the response type for Query/Params RPC method\nmessage QueryParamsResponse {\n  required Params params = 1;\n}\n')),(0,n.kt)("h2",s({},{id:"consequences"}),"Consequences"),(0,n.kt)("h3",s({},{id:"backwards-compatibility"}),"Backwards Compatibility"),(0,n.kt)("p",null,"The changes described inside this ADR are ",(0,n.kt)("strong",{parentName:"p"},"not")," backward compatible. To solve this, we will rely on the ",(0,n.kt)("inlineCode",{parentName:"p"},"x/upgrade")," module in order to properly add these new features inside a running chain. If necessary, to make sure no extra operation is performed, we should make sure that ",(0,n.kt)("inlineCode",{parentName:"p"},"fromVm[reportstypes.ModuleName]")," is set to ",(0,n.kt)("inlineCode",{parentName:"p"},"1")," before running the migrations, so that the ",(0,n.kt)("inlineCode",{parentName:"p"},"InitGenesis")," method does not get called."),(0,n.kt)("h3",s({},{id:"positive"}),"Positive"),(0,n.kt)("ul",null,(0,n.kt)("li",{parentName:"ul"},"Allow users to report misbehaving users or bad content"),(0,n.kt)("li",{parentName:"ul"},"Allow subspace owners and admins to make sure ToS are respected more easily")),(0,n.kt)("h3",s({},{id:"negative"}),"Negative"),(0,n.kt)("ul",null,(0,n.kt)("li",{parentName:"ul"},"Not known")),(0,n.kt)("h3",s({},{id:"neutral"}),"Neutral"),(0,n.kt)("ul",null,(0,n.kt)("li",{parentName:"ul"},"Require the ",(0,n.kt)("inlineCode",{parentName:"li"},"x/subspaces")," to implement the following new permissions: ",(0,n.kt)("ul",{parentName:"li"},(0,n.kt)("li",{parentName:"ul"},(0,n.kt)("inlineCode",{parentName:"li"},"PermissionCreateReport")," to allow creating new reports inside a subspace;"),(0,n.kt)("li",{parentName:"ul"},(0,n.kt)("inlineCode",{parentName:"li"},"PermissionDeleteReport")," to allow deleting existing reports inside a subspace;"),(0,n.kt)("li",{parentName:"ul"},(0,n.kt)("inlineCode",{parentName:"li"},"PermissionManageReasons")," to allow managing supported reporting reasons inside a subspace.")))),(0,n.kt)("h2",s({},{id:"further-discussions"}),"Further Discussions"),(0,n.kt)("h2",s({},{id:"references"}),"References"))}c.isMDXComponent=!0}}]);