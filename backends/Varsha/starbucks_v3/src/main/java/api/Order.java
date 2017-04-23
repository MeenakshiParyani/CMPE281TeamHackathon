package api ;

import java.util.ArrayList ;
import java.util.Random ;
import java.util.UUID ;
import java.util.concurrent.ConcurrentHashMap ;
import org.bson.types.ObjectId;

//import com.mongodb.BasicDBObject;

class Order {


	//private static final long serialVersionUID = 2105061907470199595L;
	//public String id = UUID.randomUUID().toString() ;
	
	//public String _id = new ObjectId().toString() ;
	//public String id =  new ObjectId().toString() ;
	public String _id  = new ObjectId().toString();
	public String location ; 
	public String customerName ;
    public ArrayList<OrderItem> items = new ArrayList<OrderItem>() ;
    public ConcurrentHashMap<String,String> links = new ConcurrentHashMap<String,String>();
    public StarbucksAPI.OrderStatus status ;
    public String message ;
    
}
