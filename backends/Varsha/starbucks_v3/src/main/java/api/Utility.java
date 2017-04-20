package api ;

import com.fasterxml.jackson.core.JsonProcessingException ;
import com.fasterxml.jackson.databind.ObjectMapper ;
import java.io.IOException ;
import org.bson.Document;


public class Utility {

	//Convert Document to JSON string
	public static String convertDocumentToJson ( Document doc ) {

		ObjectMapper mapper = new ObjectMapper();

		String jsonInString = null ;

		jsonInString = doc.toJson() ;
				
		return jsonInString ;

	}

	//Method to convert JSON string to Java Object
	public static Order convertJsonToOrderObj ( String jsonStr ) {

		ObjectMapper mapper = new ObjectMapper();

		String jsonInString = null ;

		Order orderObj = null ;

		try { 

				orderObj = mapper.readValue ( jsonStr , Order.class ) ;

		} catch (JsonProcessingException e) {
				// TODO Auto-generated catch block
				e.printStackTrace();
			
		}catch (IOException e) {
			
				e.printStackTrace();
		}

		return orderObj ;

	}



}