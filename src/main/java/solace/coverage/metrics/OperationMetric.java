package solace.coverage.metrics;

import java.util.Collection;
import java.util.HashSet;
import java.util.Set;
import java.util.stream.Collectors;

import io.swagger.v3.oas.models.Operation;
import io.swagger.v3.oas.models.parameters.Parameter;
import solace.coverage.annotations.InnerMetric;
import solace.coverage.base.AbstractMetric;
import solace.coverage.base.MetricMetadata;

public class OperationMetric extends AbstractMetric<Operation> {
    private Collection<Parameter> defaultParameters;
    private MetricMetadata metadata = new MetricMetadata();
    @InnerMetric
    private ParametersMetric parameters;

    @SuppressWarnings(value = "unchecked")
    public OperationMetric(Operation entity, Collection<Parameter> defaultParameters) {
        super(entity);
        this.defaultParameters = defaultParameters;
        Collection<Parameter> merged = this.mergeParameters(entity.getParameters(), defaultParameters);
        if (!merged.isEmpty())
            this.parameters = new ParametersMetric(merged);
    }

    @Override
    public float getCoverage() {
        // TODO Auto-generated method stub
        return 0;
    }

    @Override
    public void Reset() {
        this.parameters.Reset();
    }

    @Override
    public MetricMetadata getMetadata() {
        return this.metadata;
    }

    @SuppressWarnings(value = "unchecked")
    private Collection<Parameter> mergeParameters(Collection<Parameter> ...source) {
        Set<SortableParameter> parameters = new HashSet<SortableParameter>();
        for (Collection<Parameter> params : source) {
            if (params == null || params.isEmpty())
                continue;
            for (Parameter param : params)
                parameters.add(new SortableParameter(param));
        }
        return parameters.stream().map(p -> p.parameter).collect(Collectors.toSet());
    }

    private final class SortableParameter {
        private Parameter parameter;
        
        SortableParameter(Parameter parameter) {
            this.parameter = parameter;
        }
        
        @Override
        public boolean equals(Object o) {
            if (!(o instanceof Parameter))
                return false;
            else {
                Parameter p = (Parameter)o;
                return this.parameter.getName() == p.getName() && this.parameter.getIn() == p.getIn();
            }
        }

        @Override
        public int hashCode() {
            int primal = 9;
            int result = 1;
            result = primal * result + this.parameter.getName().hashCode();
            result = primal * result + this.parameter.getIn().hashCode();
            return result;
        }
    }
}
