import{o as z,$ as I,N,aw as E,ao as $,m as O,aZ as D,aM as A,C as x,L as F,a_ as U,A as j,a$ as q,B as K,b0 as y,F as V,X as L,d as Z,f as _,q as k,O as G}from"./index-PF-Fl51z.js";import{t as J}from"./el-button-C0k8b0Uu.js";import{g as Q}from"./el-select-B6CuLMMT.js";import{P as h}from"./vnode-D19UwqZF.js";const ue=(e,o,n,s)=>{let l={offsetX:0,offsetY:0};const a=u=>{const m=u.clientX,f=u.clientY,{offsetX:v,offsetY:p}=l,d=e.value.getBoundingClientRect(),g=d.left,b=d.top,S=d.width,T=d.height,Y=document.documentElement.clientWidth,B=document.documentElement.clientHeight,P=-g+v,X=-b+p,H=Y-g-S+v,W=B-b-T+p,M=C=>{let i=v+C.clientX-m,r=p+C.clientY-f;s!=null&&s.value||(i=Math.min(Math.max(i,P),H),r=Math.min(Math.max(r,X),W)),l={offsetX:i,offsetY:r},e.value&&(e.value.style.transform=`translate(${E(i)}, ${E(r)})`)},w=()=>{document.removeEventListener("mousemove",M),document.removeEventListener("mouseup",w)};document.addEventListener("mousemove",M),document.addEventListener("mouseup",w)},t=()=>{o.value&&e.value&&o.value.addEventListener("mousedown",a)},c=()=>{o.value&&e.value&&o.value.removeEventListener("mousedown",a)};z(()=>{I(()=>{n.value?t():c()})}),N(()=>{c()})},de=(e,o={})=>{$(e)||J("[useLockscreen]","You need to pass a ref param to this function");const n=o.ns||O("popup"),s=D(()=>n.bm("parent","hidden"));if(!A||x(document.body,s.value))return;let l=0,a=!1,t="0";const c=()=>{setTimeout(()=>{K(document==null?void 0:document.body,s.value),a&&document&&(document.body.style.width=t)},200)};F(e,u=>{if(!u){c();return}a=!x(document.body,s.value),a&&(t=document.body.style.width),l=Q(n.namespace.value);const m=document.documentElement.clientHeight<document.body.scrollHeight,f=U(document.body,"overflowY");l>0&&(m||f==="scroll")&&a&&(document.body.style.width=`calc(100% - ${l}px)`),j(document.body,s.value)}),q(()=>c())},R=e=>{if(!e)return{onClick:y,onMousedown:y,onMouseup:y};let o=!1,n=!1;return{onClick:t=>{o&&n&&e(t),o=n=!1},onMousedown:t=>{o=t.target===t.currentTarget},onMouseup:t=>{n=t.target===t.currentTarget}}},ee=V({mask:{type:Boolean,default:!0},customMaskEvent:{type:Boolean,default:!1},overlayClass:{type:L([String,Array,Object])},zIndex:{type:L([String,Number])}}),te={click:e=>e instanceof MouseEvent},oe="overlay";var ne=Z({name:"ElOverlay",props:ee,emits:te,setup(e,{slots:o,emit:n}){const s=O(oe),l=u=>{n("click",u)},{onClick:a,onMousedown:t,onMouseup:c}=R(e.customMaskEvent?void 0:l);return()=>e.mask?_("div",{class:[s.b(),e.overlayClass],style:{zIndex:e.zIndex},onClick:a,onMousedown:t,onMouseup:c},[k(o,"default")],h.STYLE|h.CLASS|h.PROPS,["onClick","onMouseup","onMousedown"]):G("div",{class:e.overlayClass,style:{zIndex:e.zIndex,position:"fixed",top:"0px",right:"0px",bottom:"0px",left:"0px"}},[k(o,"default")])}});const ie=ne;export{ie as E,de as a,R as b,ue as u};
