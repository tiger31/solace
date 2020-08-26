package solace.coverage.metrics;

import java.util.Map;

import io.swagger.v3.oas.models.PathItem;
import io.swagger.v3.oas.models.Operation;
import solace.coverage.base.AbstractMapMetric;
import solace.coverage.base.IRequestProcessor;
import solace.coverage.base.IResponseProcessor;
import solace.coverage.base.MetricMetadata;

public class PathItemMetric extends AbstractMapMetric<PathItem, Operation> implements IRequestProcessor, IResponseProcessor {
    private static final long serialVersionUID = 3160751773735015981L;
    private MetricMetadata metadata = new MetricMetadata();

    public PathItemMetric(PathItem entity) {
        super(entity);
        for (Map.Entry<PathItem.HttpMethod, Operation> operation : entity.readOperationsMap().entrySet()) {
            this.addMetric(operation.getKey().toString(), new OperationMetric(operation.getValue(), entity.getParameters()));
        }
    }

    @Override
    public MetricMetadata getMetadata() {
        return this.metadata;
    }
    
}