const __vite__mapDeps=(i,m=__vite__mapDeps,d=(m.f||(m.f=["assets/import-kmZ0txHV.js","assets/vendor-map-CE4d5YIX.js","assets/vendor-react-PEuRF1c_.js","assets/vendor-tdesign-Bkc5HfQ6.js","assets/modulepreload-polyfill-EV7mDuCf.js","assets/vendor-echarts-MrFpn7eE.js","assets/vendor-refine-YO-08iKi.js","assets/vendor-lib-BVPxUOb1.js","assets/vendor-tinymce-D1FYj-JL.js","assets/vendor-markdown-DBx9sqkl.js","assets/useUpload-DxbIJWgS.js","assets/export-B0u6ttO-.js","assets/useSelect-CPLYU2d6.js"])))=>i.map(i=>d[i]);
import{_ as u}from"./vendor-markdown-DBx9sqkl.js";import{j as e}from"./vendor-map-CE4d5YIX.js";import{r as f,R as x}from"./vendor-react-PEuRF1c_.js";import{A as y}from"./vendor-refine-YO-08iKi.js";import{v as h,T as j,Q as _}from"./modulepreload-polyfill-EV7mDuCf.js";import"./vendor-echarts-MrFpn7eE.js";import{d as R}from"./vendor-lib-BVPxUOb1.js";import{v as L,t as g}from"./vendor-tdesign-Bkc5HfQ6.js";import{B as b}from"./button-s0rDThSf.js";import{D as k}from"./link-Dc56DXXa.js";import"./vendor-tinymce-D1FYj-JL.js";const v=()=>{const{request:t,isLoading:r}=h();return{download:f.useCallback((a,o,c,l)=>{t(a,"post",{responseType:"blob",...c},!0).then(({response:w,...n})=>{if(n!=null&&n.data){const s=l||n.headers["content-type"],d=o||R().format("YYYYMMDD_HHmmss")+"."+j.getExtension(s);D(n.data,s,d)}else{const s=new FileReader;s.onload=function(){var m,p;const d=(p=(m=JSON.parse(s.result))==null?void 0:m.data[0])==null?void 0:p.message;L.error(d||"download error")},s.readAsText(w.data)}})},[t]),isLoading:r}},D=(t,r,i)=>{const a=new Blob([t],{type:r}),o=document.createElement("a"),l=(window.URL||window.webkitURL).createObjectURL(a);o.href=l,o.download=i,document.body.appendChild(o),o.click(),document.body.removeChild(o),window.URL.revokeObjectURL(l)},Y=()=>{const t=y(),{download:r}=v(),i=x.useMemo(()=>[{colKey:"id",sorter:!0,sortType:"all",title:"ID",width:150},{colKey:"name",title:t("tools.backup.fields.name"),ellipsis:!0},{colKey:"created_at",title:t("tools.backup.fields.createdAt"),sorter:!0,sortType:"all",width:200},{colKey:"link",title:t("table.actions"),fixed:"right",align:"center",width:150,cell:({row:a})=>e.jsxs("div",{className:"flex justify-center gap-4",children:[e.jsx(g,{theme:"primary",onClick:()=>{r(`tools/backup/download/${a.id}`)},children:"下载"}),e.jsx(k,{rowId:a.id})]})}],[t]);return e.jsx(_,{columns:i,table:{rowKey:"id"},title:t("tools.area.name"),actionRender:()=>e.jsxs(e.Fragment,{children:[e.jsx(b,{title:t("buttons.import"),component:()=>u(()=>import("./import-kmZ0txHV.js"),__vite__mapDeps([0,1,2,3,4,5,6,7,8,9,10])),action:"import",icon:e.jsx("div",{className:"t-icon i-tabler:database-import"})}),e.jsx(b,{title:t("buttons.export"),component:()=>u(()=>import("./export-B0u6ttO-.js"),__vite__mapDeps([11,1,2,3,4,5,6,7,8,9,12])),action:"export",theme:"primary",variant:"outline",icon:e.jsx("div",{className:"t-icon i-tabler:database-export"})})]})})};export{Y as default};