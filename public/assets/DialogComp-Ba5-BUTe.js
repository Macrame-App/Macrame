import { D as computed, c as createElementBlock, o as openBlock, v as renderSlot, q as normalizeClass, m as ref, b as onMounted, n as onUpdated, f as createBaseVNode, h as createVNode, w as withCtx, u as unref, E as IconX } from "./index-GNAKlyBz.js";
const _hoisted_1$1 = ["href"];
const _sfc_main$1 = {
  __name: "ButtonComp",
  props: {
    href: String,
    variant: String,
    size: String
  },
  setup(__props) {
    const props = __props;
    const classString = computed(() => {
      let classes = "btn";
      if (props.variant) classes += ` btn__${props.variant}`;
      if (props.size) classes += ` btn__${props.size}`;
      return classes;
    });
    return (_ctx, _cache) => {
      return __props.href ? (openBlock(), createElementBlock("a", {
        key: 0,
        href: __props.href,
        class: normalizeClass(classString.value)
      }, [
        renderSlot(_ctx.$slots, "default")
      ], 10, _hoisted_1$1)) : (openBlock(), createElementBlock("button", {
        key: 1,
        class: normalizeClass(classString.value)
      }, [
        renderSlot(_ctx.$slots, "default")
      ], 2));
    };
  }
};
const _hoisted_1 = { class: "dialog-container" };
const _sfc_main = {
  __name: "DialogComp",
  props: {
    open: Boolean
  },
  emits: ["onOpen", "onClose", "onToggle"],
  setup(__props, { expose: __expose, emit: __emit }) {
    const dialog = ref(null);
    const openDialog = ref();
    const emit = __emit;
    __expose({ toggleDialog });
    const props = __props;
    onMounted(() => {
      if (props.open === true) toggleDialog(props.open);
    });
    onUpdated(() => {
      if (props.open === true) toggleDialog(props.open);
    });
    function toggleDialog(openToggle) {
      if (openToggle) {
        dialog.value.showModal();
        emit("onOpen");
      } else {
        dialog.value.close();
        emit("onClose");
      }
      openDialog.value = openToggle;
      emit("onToggle");
    }
    onMounted(() => {
      openDialog.value = props.open;
      if (dialog.value.innerHTML.includes("form")) {
        dialog.value.querySelector("form").addEventListener("submit", () => {
          toggleDialog();
        });
      }
    });
    return (_ctx, _cache) => {
      return openBlock(), createElementBlock("div", _hoisted_1, [
        createBaseVNode("div", {
          class: "trigger",
          onClick: _cache[0] || (_cache[0] = ($event) => toggleDialog(true))
        }, [
          renderSlot(_ctx.$slots, "trigger")
        ]),
        createBaseVNode("dialog", {
          ref_key: "dialog",
          ref: dialog,
          class: "mcrm-block block__dark"
        }, [
          createVNode(_sfc_main$1, {
            class: "dialog__close p-0",
            variant: "ghost",
            size: "sm",
            tabindex: "-1",
            onClick: _cache[1] || (_cache[1] = ($event) => toggleDialog(false))
          }, {
            default: withCtx(() => [
              createVNode(unref(IconX))
            ]),
            _: 1
          }),
          renderSlot(_ctx.$slots, "content")
        ], 512)
      ]);
    };
  }
};
export {
  _sfc_main$1 as _,
  _sfc_main as a
};
