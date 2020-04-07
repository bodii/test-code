import mongoose from 'mongoose';
import ItemsSchema from '../../schemas/surname/items';


mongoose.connect('mongodb://localhost/surname');

const Items = mongoose.model('Items', ItemsSchema);


export default Items;