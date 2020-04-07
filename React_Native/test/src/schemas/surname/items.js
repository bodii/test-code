import { Schema } from "mongoose";

const infoSchema = new Schema({
  拼音: {
    type: String,
    default: "",
  },
  郡望: {
    type: String,
    default: "",
  },
  名人: {
    type: String,
    default: "",
  },
  胜迹: {
    type: String,
    default: "",
  },
  文献: {
    type: String,
    default: "",
  },
  历史: {
    type: String,
    default: "",
  },
});

const ItemsSchema = new Schema(
  {
    name: {
      type: String,
      default: "",
      required: true,
    },
    title: {
      type: String,
      default: "",
    },
    url: {
      type: String,
      default: "",
    },
    belong_to: {
      type: String,
      default: "",
    },
    info: {
      type: infoSchema,
      default: "",
    },
  },
  { collection: "item" }
);

export default ItemsSchema;
