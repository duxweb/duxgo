const __vite__mapDeps=(i,m=__vite__mapDeps,d=(m.f||(m.f=["assets/group-BW3iEmvv.js","assets/vendor-map-CE4d5YIX.js","assets/vendor-react-PEuRF1c_.js","assets/vendor-tdesign-Bkc5HfQ6.js","assets/vendor-refine-YO-08iKi.js","assets/modulepreload-polyfill-EV7mDuCf.js","assets/vendor-echarts-MrFpn7eE.js","assets/vendor-lib-BVPxUOb1.js","assets/vendor-tinymce-D1FYj-JL.js","assets/vendor-markdown-DBx9sqkl.js","assets/upload-O3lVDCuc.js","assets/uploadFile-YsBOGyhm.js","assets/useUpload-DxbIJWgS.js"])))=>i.map(i=>d[i]);
import{_ as a}from"./vendor-markdown-DBx9sqkl.js";import{j as e}from"./vendor-map-CE4d5YIX.js";import{r as c,R as o}from"./vendor-react-PEuRF1c_.js";import{A as m}from"./vendor-refine-YO-08iKi.js";import{Q as f}from"./modulepreload-polyfill-EV7mDuCf.js";import"./vendor-echarts-MrFpn7eE.js";import"./vendor-lib-BVPxUOb1.js";import{B as x}from"./button-s0rDThSf.js";import{D as p}from"./link-Dc56DXXa.js";import{F as h}from"./filterSider-DROQ4EVH.js";import"./vendor-tinymce-D1FYj-JL.js";import{t as u}from"./vendor-tdesign-Bkc5HfQ6.js";const j=({mime:t})=>{switch(!0){case/^image\//.test(t):return e.jsx("div",{className:"h-10 w-10 flex items-center justify-center rounded p-2 text-white bg-brand",children:e.jsx("div",{className:"i-tabler:photo h-6 w-6"})});case/^video\//.test(t):return e.jsx("div",{className:"h-10 w-10 flex items-center justify-center rounded p-2 text-white bg-success",children:e.jsx("div",{className:"i-tabler:video h-6 w-6"})});case/^audio\//.test(t):return e.jsx("div",{className:"h-10 w-10 flex items-center justify-center rounded p-2 text-white bg-warning",children:e.jsx("div",{className:"i-tabler:audio h-6 w-6"})});case/^application\/pdf$/.test(t):return e.jsx("div",{className:"h-10 w-10 flex items-center justify-center rounded p-2 text-white bg-error",children:e.jsx("div",{className:"i-tabler:file-pdf h-6 w-6"})});case/^application\/vnd\.openxmlformats-officedocument\.wordprocessingml\.document$/.test(t):case/^application\/msword$/.test(t):return e.jsx("div",{className:"h-10 w-10 flex items-center justify-center rounded p-2 text-white bg-brand",children:e.jsx("div",{className:"i-tabler:file-word h-6 w-6"})});case/^application\/vnd\.openxmlformats-officedocument\.spreadsheetml\.sheet$/.test(t):case/^application\/vnd\.ms-excel$/.test(t):return e.jsx("div",{className:"h-10 w-10 flex items-center justify-center rounded p-2 text-white bg-brand",children:e.jsx("div",{className:"i-tabler:file-excel h-6 w-6"})});case/^application\/zip$/.test(t):case/^application\/x-rar-compressed$/.test(t):case/^application\/x-7z-compressed$/.test(t):return e.jsx("div",{className:"h-10 w-10 flex items-center justify-center rounded p-2 text-white bg-brand",children:e.jsx("div",{className:"i-tabler:file-zip h-6 w-6"})});case/^text\//.test(t):return e.jsx("div",{className:"h-10 w-10 flex items-center justify-center rounded p-2 text-white bg-brand",children:e.jsx("div",{className:"i-tabler:file-text h-6 w-6"})});default:return e.jsx("div",{className:"h-10 w-10 flex items-center justify-center rounded p-2 text-white bg-brand",children:e.jsx("div",{className:"i-tabler:file-unknown h-6 w-6"})})}},L=()=>{var r,l;const t=m(),s=c.useRef(null),d=o.useMemo(()=>[{colKey:"name",title:t("tools.file.fields.name"),minWidth:300,cell:({row:i})=>e.jsxs("div",{className:"flex items-center gap-2",children:[e.jsx("div",{children:e.jsx(j,{mime:i.mime})}),e.jsxs("div",{className:"flex flex-col",children:[e.jsx("div",{children:i.name}),e.jsx("div",{className:"text-sm text-gray",children:i.mime})]})]})},{colKey:"dir_name",title:t("tools.file.fields.dir"),width:150},{colKey:"size",title:t("tools.file.fields.size"),width:150},{colKey:"driver",title:t("tools.file.fields.driver"),width:150},{colKey:"time",title:t("tools.file.fields.time"),width:200},{colKey:"link",title:t("table.actions"),fixed:"right",align:"center",width:160,cell:({row:i})=>e.jsxs("div",{className:"flex justify-center gap-4",children:[e.jsx(u,{theme:"primary",href:i.url,target:"_block",children:t("buttons.show")}),e.jsx(p,{rowId:i.id})]})}],[t]),n=(l=(r=s.current)==null?void 0:r.filters)==null?void 0:l.dir_id;return e.jsx(e.Fragment,{children:e.jsx(f,{ref:s,columns:d,table:{rowKey:"id"},siderRender:()=>e.jsx(h,{title:t("tools.file.fields.dir"),component:()=>a(()=>import("./group-BW3iEmvv.js"),__vite__mapDeps([0,1,2,3,4,5,6,7,8,9])),resource:"tools.fileDir",field:"dir_id",optionLabel:"name",optionValue:"id"}),actionRender:()=>e.jsx(x,{component:()=>a(()=>import("./upload-O3lVDCuc.js"),__vite__mapDeps([10,1,2,3,4,5,6,7,8,9,11,12])),action:"upload",title:t("tools.file.fields.upload"),icon:e.jsx("div",{className:"t-icon i-tabler:plus"}),rowId:n})})})};export{L as default};