
package api ;

import java.util.concurrent.BlockingQueue;

public class StarbucksBarista implements Runnable {

	protected BlockingQueue<String> blockingQueue ;

	public StarbucksBarista(BlockingQueue<String> queue) {
		this.blockingQueue = queue;
	}

	@Override
	public void run() {
		System.out.println( "Barista thread started!" ) ;
		while (true) {
			try {
				try { 
					System.out.println("In Barista : GOing to Sleep !!");
					Thread.sleep(10000); 
					System.out.println("In Barista : Woke up !!");
				} catch ( Exception e ) {

					System.out.println("In Barista : Exception!! ");
					e.printStackTrace();

				}
				System.out.println("In Barista : Picking up next order!!!" ) ;  
				String order_id = blockingQueue.take();
				System.out.println("In Barista : Order ID : " + order_id) ;
				Order order = StarbucksAPI.getOrder( order_id ) ;
				System.out.println("In Barista : Got order for process") ;
        		if ( order != null && order.status == StarbucksAPI.OrderStatus.PAID ) {
					System.out.println(Thread.currentThread().getName() + " Processed Order: " + order_id );
					StarbucksAPI.setOrderStatus( order, StarbucksAPI.OrderStatus.PREPARING ) ;
					StarbucksAPI.updateOrder( order._id, order ) ; 
					try { Thread.sleep(20000); } catch ( Exception e ) {}  
					StarbucksAPI.setOrderStatus( order, StarbucksAPI.OrderStatus.SERVED ) ;					
					StarbucksAPI.updateOrder( order._id, order ) ; 
					try { Thread.sleep(10000); } catch ( Exception e ) {}  
					StarbucksAPI.setOrderStatus( order, StarbucksAPI.OrderStatus.COLLECTED ) ;
					StarbucksAPI.updateOrder( order._id, order ) ; 				
				}
				else {
					System.out.println( "Order is null or the order is not Paid" ) ;
					blockingQueue.put( order_id ) ;
				}
			} catch (InterruptedException e) {
				e.printStackTrace();
			}
		}
	}

}