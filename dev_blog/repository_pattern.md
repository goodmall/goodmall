http://codebetter.com/gregyoung/2009/04/23/repository-is-dead-long-live-repository/

>
	Provides the domain collection semantics to a set of aggregate root objects.

	So now that we have determined that he has not actually come up with anything new and is actually still using repositories let’s reframe the original argument into what it really is.
	
	FindCustomer(id)
	FindCustomerWithAddresses(id)
	FindCustomerWith..
	Oren has an issue with named query methods on a Repository interface. I will ignore the arguments provided by him and provide a simpler better argument for why you SHOULD NOT use these types of methods.
	
	The Open Closed Principle. Versioning a Repository that uses these types of methods is extremely difficult. When I want to add a new method, or change an existing one I am changing the original contract. If I move these to a single method that takes a Query object I can change/add/remove queries without touching my Repository contract. In other words, we are simply applying SOLID to the repository interface.
	Now we all know that following the Open Closed Principle amongst others is good right?

	I would answer that question with “Yes in general following the SOLID principles leads to better Object Oriented code”.
	
	I am sure I have left a few fairly confused at this point as I argue against my own point but I am leading to something important. The problem here is that the Repository interface is not necessarily Object Oriented. The Repository represents an architectural boundary, it is intended to be a LAYER/TIER boundary. Generally speaking when we define such interfaces we define them in a procedural manner (and with good cause).
	
	Analyzing the situation given of a CustomerRepository what would happen if we were to want to put the data access behind a remote facade? With the simple procedural boundary of named methods, we would just go create a remote facade (say a webservice) and we would pass through the calls. What would happen though if we used the “other” Repository interface that is being suggested? Well our remote facade would need to support the passing of any criteria dynamically across its contract, this is generally considered bad contract design as we then will have great trouble figuring out and optimizing what our service actually does. With the explicit contract we have to explicitly add or change the contract and we know when things are being added that may need to be optimized. I am quite sure everyone has dealt at some point with the service that had 1 method and could do anything typically they take and receive strings my favorite are the ones you pass SQL to and they return you random XML (usually a serialized data table) …
	
	One could argue against me here by saying that they don’t consider their Repositories to be LAYER/TIER boundaries… Well, I am sorry but you are using the pattern incorrectly then. If you don’t want a LAYER/TIER boundary don’t have one just use nhibernate directly … At this point you probably shouldn’t have a domain either though … If your system is complex enough to justify the cost of creating a good model why are you trying to optimize the < 2% spent on Repository abstractions?
	
	 
	
	After this Oren brings forth another very interesting issue that led him to his epiphany of killing Repositories.
	
	It get worse when you have complex search criteria and complex fetch plan. Then you are stuck either creating a method per each combination that you use or generalizing that
	
	I have had this smell in the past as well but instead of destroying the layering I am building into my domain (with good reason, see DDD by Evans for why) I went a completely different route. I noticed very quickly that it was by some random chance that my fetch plans were being different. I had a very distinct place where things were different, I needed very different fetching plans between when I was getting domain objects to perform a writing behaviour on them as opposed to when I was reading objects to say build a DTO.
	
	This realization is what led me to command and query separation, the creation of a separate layer to process read->DTO transformations and to no longer use my domain repositories for this purpose. I should point out that for the read layer, what is being shown by Oren is a great way of implementing it. I won’t go into a large talk on command and query separation here (there is a good video from the european van and hopefully the QCon SF session will be up soon)… but if you apply command and query separation you will have almost none (read: none) read methods on your repositories in your domain.
	
	Continuing along, Oren randomly puts forth the gem of:
	
	A lot of people use it, mostly because of the DDD association. I am currently in the opinion that DDD should be approached with caution, since if you don’t actually need it (and have the prerequisites for it, such as business expert to work closely with or an app that can actually benefit from it), it is probably going to be more painful to try using DDD than without.
	
	It being everyone has heard pretty much every real DDD practitioner saying this for years, I am really unsure why it has taken so long to figure this out?  This also begs the question of what systems actually need or are suitable for DDD but that’s another day and another post.

## from http://debasishg.blogspot.com/2007/02/domain-driven-design-inject.html

Domain Driven Design - Inject Repositories, not DAOs in Domain Entities	

> Regarding "having fine-grained methods withing DAO", yes, I tend to use DAO as a lower level abstraction and build methods of higher granularity (closer to the domain) at the Repository. You are correct .. this may lead to some inefficiency since I could have used one performant SQL instead of multiple queries. This is all a compromise between modularity and performance. I start with a focus on the former and then when some of the areas need blazing performance I fall down to optimize and club them into a single SQL. Remember not all parts of your application need blazing performance .. for the ones that need it, you can compromise on the modularity.	