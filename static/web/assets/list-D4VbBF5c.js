const __vite__mapDeps=(i,m=__vite__mapDeps,d=(m.f||(m.f=["assets/show-Dm7QgAqv.js","assets/vendor-map-CE4d5YIX.js","assets/vendor-react-PEuRF1c_.js","assets/vendor-tdesign-Bkc5HfQ6.js","assets/vendor-refine-YO-08iKi.js","assets/modulepreload-polyfill-EV7mDuCf.js","assets/vendor-echarts-MrFpn7eE.js","assets/vendor-lib-BVPxUOb1.js","assets/vendor-tinymce-D1FYj-JL.js","assets/vendor-markdown-DBx9sqkl.js"])))=>i.map(i=>d[i]);
import{_ as d}from"./vendor-markdown-DBx9sqkl.js";import{j as e}from"./vendor-map-CE4d5YIX.js";import{R as p}from"./vendor-react-PEuRF1c_.js";import{A as h}from"./vendor-refine-YO-08iKi.js";import{S as r,Q as u,F as l}from"./modulepreload-polyfill-EV7mDuCf.js";import"./vendor-echarts-MrFpn7eE.js";import"./vendor-lib-BVPxUOb1.js";import{S as x}from"./link-Dc56DXXa.js";import"./vendor-tinymce-D1FYj-JL.js";import{u as j}from"./useSelect-CPLYU2d6.js";import{a5 as f,t as y,X as a,S as i,a3 as v}from"./vendor-tdesign-Bkc5HfQ6.js";const N=()=>{const s=h(),n=p.useMemo(()=>[{colKey:"id",sorter:!0,sortType:"all",title:"ID",width:100},{colKey:"username",title:s("system.operate.fields.user"),cell:({row:t})=>e.jsxs(r,{size:"small",children:[e.jsx(r.Avatar,{image:t.avatar,children:t.nickname}),e.jsx(r.Title,{children:t.nickname}),e.jsx(r.Desc,{children:t.username})]})},{colKey:"request_method",title:s("system.operate.fields.request"),minWidth:210,cell:({row:t})=>e.jsxs("div",{className:"flex flex-col gap-2",children:[e.jsx("div",{children:e.jsx(f,{content:t.request_url,children:e.jsx(y,{children:t.route_name})})}),e.jsxs("div",{className:"flex gap-2",children:[e.jsx(a,{theme:"primary",variant:"outline",children:t.request_method}),e.jsx(a,{theme:"success",variant:"outline",children:t.request_time})]})]})},{colKey:"client_ip",title:s("system.operate.fields.client"),minWidth:200,cell:({row:t})=>e.jsxs("div",{className:"flex flex-col gap-2",children:[e.jsx("div",{children:t.client_ip}),e.jsx("div",{className:"flex gap-2",children:e.jsx(a,{theme:"primary",variant:"outline",children:t.client_device})})]})},{colKey:"time",title:s("system.operate.fields.requestTime"),minWidth:200},{colKey:"link",title:s("table.actions"),fixed:"right",align:"center",width:160,cell:({row:t})=>e.jsx("div",{className:"flex justify-center gap-4",children:e.jsx(x,{component:()=>d(()=>import("./show-Dm7QgAqv.js"),__vite__mapDeps([0,1,2,3,4,5,6,7,8,9])),rowId:t.id})})}],[s]),{options:o,onSearch:c,queryResult:m}=j({resource:"system.user",optionLabel:"nickname",optionValue:"id"});return e.jsx(u,{columns:n,table:{rowKey:"id"},title:s("system.operate.name"),filterRender:()=>e.jsxs(e.Fragment,{children:[e.jsx(l,{name:"user",children:e.jsx(i,{filterable:!0,loading:m.isLoading,onSearch:c,options:o,placeholder:s("system.operate.filters.userPlaceholder"),clearable:!0})}),e.jsx(l,{name:"method",children:e.jsxs(i,{placeholder:s("system.operate.filters.method.placeholder"),clearable:!0,children:[e.jsx(i.Option,{value:"post",children:s("system.operate.filters.method.post")}),e.jsx(i.Option,{value:"put",children:s("system.operate.filters.method.put")}),e.jsx(i.Option,{value:"patch",children:s("system.operate.filters.method.patch")}),e.jsx(i.Option,{value:"delete",children:s("system.operate.filters.method.delete")})]})}),e.jsx(l,{name:"date",children:e.jsx(v,{enableTimePicker:!0,clearable:!0})})]})})};export{N as default};