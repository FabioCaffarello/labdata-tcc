from dataclasses import dataclass, field
from typing import Dict, Any


@dataclass
class ErrMsgDTO:
    # Error message
    err: Exception = field(metadata={"json": "error"})
    # Listener tag
    listener_tag: str = field(metadata={"json": "listener_tag"})
    # Message
    msg: bytes = field(metadata={"json": "msg"})


@dataclass
class ProcessOrderDTO:
    # ID represents the unique identifier of the order.
    order_id: str = field(metadata={"json": "_id"})
    # ProcessingID represents the unique identifier of the order processing.
    processing_id: str = field(metadata={"json": "processing_id"})
    # Service represents the name of the service for which the order is processed.
    service: str = field(metadata={"json": "service"})
    # Source indicates the origin or source of the order.
    source: str = field(metadata={"json": "source"})
    # Provider specifies the provider of the order.
    provider: str = field(metadata={"json": "provider"})
    # Stage represents the current stage of the order processing.
    stage: str = field(metadata={"json": "stage"})
    # InputID represents the unique identifier of the input data.
    input_id: str = field(metadata={"json": "input_id"})
    # Data represents the order data.
    data: Dict[str, Any] = field(metadata={"json": "data"})
