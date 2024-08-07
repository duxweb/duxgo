import{j as e}from"./vendor-map-CE4d5YIX.js";import{u as q,h as A}from"./useUpload-DxbIJWgS.js";import{L as F,I as L,J as U,O as P,K as S,P as B,R as E,x as H,v as O,j as Y}from"./modulepreload-polyfill-EV7mDuCf.js";import{r as g}from"./vendor-react-PEuRF1c_.js";import{d as T}from"./vendor-lib-BVPxUOb1.js";import{A as V}from"./vendor-refine-YO-08iKi.js";import{U as I,t as _,B as M,a8 as G}from"./vendor-tdesign-Bkc5HfQ6.js";function J(r,c,o,t){return r.config&&r.config.forceDirect?(t.info("ues forceDirect mode."),new S(r,c,o,t)):r.file.size>4*B?(t.info("file size over 4M, use Resume."),new E(r,c,o,t)):(t.info("file size less or equal than 4M, use Direct."),new S(r,c,o,t))}function K(r,c,o,t,i){var h=new F(o,i==null?void 0:i.disableStatisticsReport,i==null?void 0:i.debugLogLevel,r.name),p={file:r,key:c,token:o,putExtra:t,config:L(i,h)},w=new U(p.config.uphost);return new P(function(f){var m=J(p,{onData:function(n){return f.next(n)},onError:function(n){return f.error(n)},onComplete:function(n){return f.complete(n)}},w,h);return m.putFile(),m.stop.bind(m)})}const ne=({className:r,value:c,defaultValue:o,onChange:t,multiple:i,hookProps:h,driver:p,...w})=>{const f=V(),m=q(h),[n,j]=H({value:c,defaultValue:o,onChange:t}),{request:y}=O(),N=g.useRef(null),C=g.useMemo(()=>n?Array.isArray(n)?n:[n]:[],[n]),D=g.useCallback(l=>new Promise(s=>{let a=l;Array.isArray(l)&&(a=l[0]);const v=Q((a==null?void 0:a.name)||"");y("upload/qiniu","POST").then(d=>{var z,k;K(a.raw,v,(z=d==null?void 0:d.data)==null?void 0:z.token,{fname:a.name},{uphost:(k=d==null?void 0:d.data)==null?void 0:k.domain}).subscribe({next({total:u}){var x;(x=N.current)==null||x.uploadFilePercent({file:a,percent:Math.round(u==null?void 0:u.percent)||0})},error(u){s({status:"fail",response:{url:"",error:u.message}})},complete({key:u}){var x;s({status:"success",response:{url:((x=d==null?void 0:d.data)==null?void 0:x.public_url)+"/"+u,size:a.size,name:a.name}})}})})}),[y,N]),[R,b]=g.useState(!1);return e.jsx(I,{ref:N,className:Y([r,"w-full app-upload-file"]),files:C,onChange:l=>{const s=l==null?void 0:l.map(a=>({url:a.url,name:a.name,size:a.size,uploadTime:a.uploadTime,status:"success"}));j(i?s:s==null?void 0:s[0])},onSelectChange:()=>{b(!0)},onSuccess:()=>{b(!1)},onValidate:()=>{b(!1)},onFail:()=>{b(!1)},multiple:i,...m,...w,draggable:!1,theme:"file",requestMethod:p=="qiniu"?D:void 0,formatResponse:p=="qiniu"?void 0:m.formatResponse,fileListDisplay:({files:l})=>e.jsx("div",{className:"mt-4 flex flex-col gap-2",children:l.map((s,a)=>e.jsxs("div",{className:"border rounded py-2 pl-4 pr-2 border-component",children:[e.jsxs("div",{className:"flex items-center justify-between",children:[e.jsxs("div",{className:"flex flex-col",children:[e.jsxs("div",{className:"flex gap-2 items-center",children:[s.status=="success"&&e.jsx("div",{className:"bg-success text-white rounded-full flex items-center justify-center p-0.5",children:e.jsx("div",{className:"i-tabler:check w-2.5 h-2.5"})}),s.status=="fail"&&e.jsx("div",{className:"bg-error text-white rounded-full flex items-center justify-center p-0.5",children:e.jsx("div",{className:"i-tabler:x w-2.5 h-2.5"})}),e.jsx(_,{href:s.url,target:"_blank",children:s.name})]}),e.jsx("div",{className:"text-xs text-placeholder",children:A(s.size||0)})]}),e.jsxs("div",{className:"flex",children:[s.url&&e.jsx(M,{theme:"default",variant:"text",shape:"circle",icon:e.jsx("div",{className:"i-tabler:download"}),onClick:()=>{window.open(s.url)}}),!(s!=null&&s.status)||s.status=="success"?e.jsx(M,{theme:"default",variant:"text",shape:"circle",icon:e.jsx("div",{className:"i-tabler:x"}),onClick:()=>{if(Array.isArray(l)){const v=[...l];v.splice(a,1),j(v)}else j(void 0)}}):""]})]}),s.status=="progress"&&e.jsx("div",{children:e.jsx(G,{status:W(s.status),percentage:s.percent,theme:"line"})})]},a))}),children:e.jsxs("div",{className:"w-full flex flex-col items-center rounded bg-gray-1 p-6 dark:bg-gray-12",children:[e.jsx("div",{className:"",children:e.jsxs("svg",{className:"h-12 w-12",viewBox:"0 0 1024 1024",version:"1.1",xmlns:"http://www.w3.org/2000/svg",width:"64",height:"64",children:[e.jsx("path",{d:"M864.19 493.58c85.24 8.4 158.5 87.38 158.5 191.92 0 91.49-82.02 173.51-173.51 173.51H213.21v-0.18C-41 840.95-70 518.33 147.45 438 129.33 309.59 255.57 217.97 367.6 264.14c196.83-214.92 540.45-53.91 496.59 229.44z m0 0",className:"fill-brand-7 dark:fill-brand-8"}),e.jsx("path",{d:"M867.22 493.94c-1.07-0.18-2.14-0.18-3.22-0.36 43.14-295.65-310.96-437.55-496.59-229.62-118.36-47.85-237.84 51.62-220.15 173.87-224.38 85.51-180.37 405.97 65.77 420.82v0.18H586.5c144.38-68.27 251.06-203.18 280.72-364.89z m0 0",className:"fill-brand-6 dark:fill-brand-9"}),e.jsx("path",{d:"M6.29 598.47C22.55 522.89 75.97 463.02 147.45 438 129.33 309.59 255.57 217.97 367.6 264.14c0.72-0.72 1.43-1.61 1.96-2.32-53.6 166.72-191.37 295.55-363.27 336.65z m0 0",className:"fill-brand-7 dark:fill-brand-8"}),e.jsx("path",{d:"M581.85 859.01H446.94v-221.4c-134.72 6.06-84.29-25.48-20.26-95.29 50.55-49.97 81.67-101.29 101.54-85.2 150.45 167.98 189.27 185.36 53.46 180.3V859h0.17z m0 0",className:"fill-white/80"})]})}),e.jsx("div",{className:"text-placeholder mt-2",children:f("fields.placeholder",{ns:"file"})}),e.jsx("div",{className:"flex mt-4",children:e.jsx(M,{icon:e.jsx("div",{className:"i-tabler:upload t-icon"}),loading:R,children:f("fields.file",{ns:"file"})})})]})})},Q=r=>{const c=r.substring(r.lastIndexOf(".")),t=Math.random().toString(36).substring(2,15)+Math.random().toString(36).substring(2,15)+c;return T().format("YYYY-MM-DD")+"/"+t},W=r=>{switch(r){case"success":return"success";case"progress":return"active";case"waiting":return"warning";case"error":return"error"}};export{ne as U};