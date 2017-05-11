package api;


import com.mongodb.MongoClient;
import com.mongodb.MongoClientURI;
import com.mongodb.ServerAddress;

import com.mongodb.client.MongoDatabase;
import com.mongodb.client.MongoCollection;
import com.mongodb.DBObject;
import org.bson.types.ObjectId;
import com.mongodb.util.JSON;

import org.bson.Document;
import java.util.Arrays;
import com.mongodb.Block;

import com.mongodb.client.MongoCursor;
import static com.mongodb.client.model.Filters.*;
import com.mongodb.client.result.DeleteResult;
import static com.mongodb.client.model.Updates.*;
import com.mongodb.client.result.UpdateResult;
import java.util.ArrayList;
import java.util.List;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;

import com.mongodb.BasicDBObject;

import com.mongodb.client.model.UpdateOptions;

import java.io.IOException;
import java.util.Collection ;
import java.util.ArrayList ;



public class DbOperations { 


	private static MongoClient mongoClient = new MongoClient( "54.219.173.227" , 27017 )  ;


	private static MongoCollection<Document> getCollection ( String collName ) {


		MongoDatabase database = mongoClient.getDatabase("starbucksv3") ;

		return database.getCollection( collName ) ;


	}



	public static void insertOrder ( String order_id , Order order )  {

		MongoCollection<Document> orderColl  =  getCollection( "order" );

		ObjectMapper mapper = new ObjectMapper();
		try {
			
			String jsonInString = mapper.writeValueAsString(order);

			orderColl.insertOne( Document.parse( jsonInString ) );

		} catch (JsonProcessingException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		}
		
	}


	public static void updateOrder ( String order_id , Order order ) {

		MongoCollection<Document> orderColl  =  getCollection( "order" );

		ObjectMapper mapper = new ObjectMapper();
		
		try {
		
			String jsonInString = mapper.writeValueAsString(order);
		
			orderColl.replaceOne (  
				eq ( "_id" , order_id ) , 
				Document.parse ( jsonInString ) , 
				new UpdateOptions ().upsert ( false ) .bypassDocumentValidation ( true ) 
			) ;

		} catch (JsonProcessingException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		}



	}



	public static Order getOrder ( String order_id  ) {


		ObjectMapper mapper = new ObjectMapper();

		Order orderObj = null ;

		MongoCollection<Document> orderColl  =  getCollection( "order" );
		
		//1. Find the record in db and get it as documnet

		Document orderDtlsDoc = orderColl.find ( eq ( "_id" , order_id ) ).first() ;

		//System.out.println("DbOperations : Getting order from DB !! " + orderDtlsDoc.getString("location"));

		//2. Convert the Document to JSON string
		if(orderDtlsDoc != null ) {

			System.out.println("DbOperations : Order not null !! ");

			String orderJsonStr = orderDtlsDoc.toJson();

			System.out.println( "DbOperations  : String Json : " + orderJsonStr ) ; 

			//3. Convert the JSON string into Order object
			try { 

				orderObj = mapper.readValue ( orderJsonStr , Order.class ) ;

				System.out.println( "DBOperations  :  Order object after conversion : " + orderObj.location ) ; 

			} catch (JsonProcessingException e) {
				// TODO Auto-generated catch block
				e.printStackTrace();
			
			}catch (IOException e) {
			
				e.printStackTrace();
			}

		}else{

			System.out.println("---------In DBOperations : Null document for order ") ;
		}

		return orderObj ;

	}


	public static void removeOrder ( String order_id ) {

		MongoCollection<Document> orderColl  =  getCollection( "order" );

		orderColl.deleteOne ( eq ( "_id" , order_id ) ) ;

		System.out.println( "Order deleted ------- " ) ;

	}



	public static Collection<Order> getOrders () {

		MongoCollection<Document> orderColl  =  getCollection( "order" );

		List<Document> orderDocList = (List<Document>) orderColl.find().into(
				new ArrayList<Document>());


		ArrayList<Order> orderList = new ArrayList<Order>() ;

		for (Document orderDoc  : orderDocList ) {

				String jsonInString = Utility.convertDocumentToJson ( orderDoc ) ;
				orderList.add ( Utility.convertJsonToOrderObj ( jsonInString ) ) ;

		}
 
		return orderList ;

	}

}
